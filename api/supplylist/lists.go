package supplylist

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/api/supplies"
	"school-supply-list/database"
	"strconv"
)

type SupplyList struct {
	ListID              int                              `json:"list_id"`
	Grade               int                              `json:"grade"`
	SchoolID            int                              `json:"school_id"`
	ListName            string                           `json:"list_name"`
	BasicSupplies       []supplies.SupplyItem            `json:"basic_supplies"`
	CategorizedSupplies map[string][]supplies.SupplyItem `json:"categorized_supplies"`
	Published           bool                             `json:"published"`
	Checked             []int                            `json:"checked"`
}

func CreateSupplyList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var list SupplyList
		err := c.BindJSON(&list)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}
		row := db.Db.QueryRow(`INSERT INTO supply_list (grade, list_name, school_id, published, list_id) 
		  VALUES ($1, $2, $3, false, default) returning list_id`, list.Grade, list.ListName, list.SchoolID)
		err = row.Scan(&list.ListID)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, list)
	}
}

func GetSupplyList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		list := SupplyList{ListID: id}
		rowCount := -1
		rows, err := db.Db.Query(`SELECT list_id, grade, list_name, school_id from supply_list 
											where supply_list.list_id=$1`, id)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		for rows.Next(){
			err = rows.Scan(&list.ListID, &list.Grade, &list.ListName, &list.SchoolID)
			if err != nil {
				database.CheckDBErr(err.(*pq.Error), c)
				return
			}
			rowCount ++
		}

		if rowCount == -1{
			c.AbortWithStatusJSON(404, "This resource does not exist")
			return
		}

		session, err := db.SessionStore.Get(c.Request, "session")
		if err != nil {
			c.AbortWithStatusJSON(500, "The server was unable to retrieve this session")
			return
		}
		googleID := session.Values["GoogleId"]
		if googleID == nil{
			list.BasicSupplies, list.CategorizedSupplies, err = getItemsForList(list.ListID, db)
		}else{
			list.BasicSupplies, list.CategorizedSupplies, list.Checked, err = getItemsForUserList(list.ListID, db, googleID.(string))
		}
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, list)
	}
}

func getItemsForList(id int, db *database.DB) ([]supplies.SupplyItem, map[string][]supplies.SupplyItem, error) {
	var basicSupplies []supplies.SupplyItem
	rows, err := db.Db.Query(`SELECT id, list_id, supply_name, supply_desc, category FROM supply_item sup 
										WHERE list_id = $1`, id)
	categorizedSupplies := make(map[string][]supplies.SupplyItem)
	if err != nil {
		return basicSupplies, categorizedSupplies, err
	}
	for rows.Next() {
		var supply supplies.SupplyItem

		err = rows.Scan(&supply.ID, &supply.ListID, &supply.Supply, &supply.Desc, &supply.Category)
		if err != nil {
			return basicSupplies, categorizedSupplies, err
		}
		// Check if item is categorized
		if supply.Category.Valid {
			// Either create a new map item or add to existing item
			if val, ok := categorizedSupplies[supply.Category.String]; ok {
				val = append(val, supply)
			} else {
				categorizedSupplies[supply.Category.String] = []supplies.SupplyItem{supply}
			}
		} else {
			basicSupplies = append(basicSupplies, supply)
		}
	}
	return basicSupplies, categorizedSupplies, nil
}

func getItemsForUserList(id int, db *database.DB, googleID string) ([]supplies.SupplyItem, map[string][]supplies.SupplyItem, []int, error) {
	var basicSupplies []supplies.SupplyItem
	categorizedSupplies := make(map[string][]supplies.SupplyItem)
	var checked []int
	var userID string
	userRows := db.Db.QueryRow(`SELECT user_id from account a where a.google_id=$1`, googleID)
	if userRows.Err() != nil {
		return basicSupplies, categorizedSupplies, checked, userRows.Err()
	}
	err := userRows.Scan(&userID)
	if err != nil {
		return basicSupplies, categorizedSupplies, checked, err
	}

	rows, err := db.Db.Query(`SELECT id, list_id, supply_name, supply_desc, category,
       									CASE WHEN user_uuid IS NOT NULL
       									    THEN 'true' ELSE 'false' END FROM supply_item sup
										JOIN checked_items ci on sup.id = ci.item_id
										WHERE list_id = $1 AND user_uuid=$2`, id, userID)
	if err != nil {
		return basicSupplies, categorizedSupplies, checked, err
	}
	for rows.Next() {
		var supply supplies.SupplyItem
		var isChecked bool
		err = rows.Scan(&supply.ID, &supply.ListID, &supply.Supply, &supply.Desc, &supply.Category, &isChecked)

		// Check if item is categorized
		if supply.Category.Valid {
			// Either create a new map item or add to existing item
			if val, ok := categorizedSupplies[supply.Category.String]; ok {
				val = append(val, supply)
			} else {
				categorizedSupplies[supply.Category.String] = []supplies.SupplyItem{supply}
			}
		} else {
			basicSupplies = append(basicSupplies, supply)
		}

		if isChecked{
			checked = append(checked, supply.ID)
		}
	}
	return basicSupplies, categorizedSupplies, checked, nil
}

func UpdateSupplyList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		var list SupplyList
		err = json.NewDecoder(c.Request.Body).Decode(&list)

		list.ListID = id

		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}
		row := db.Db.QueryRow(`UPDATE supply_list SET grade=$1, list_name=$2, published=$3, 
	   		school_id=$4 where list_id=$5 returning list_id`, list.Grade, list.ListName, list.Published, list.SchoolID, list.ListID)
		err = row.Scan(&list.ListID)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, list)
	}
}

func UpdateSavedList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		listID, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid listID. Must be an integer.")
			return
		}
		var checked []int
		err = json.NewDecoder(c.Request.Body).Decode(&checked)

		session, err := db.SessionStore.Get(c.Request, "session")
		if err != nil {
			c.AbortWithStatusJSON(500, "The server was unable to retrieve this session")
			return
		}

		googleID := session.Values["GoogleId"]
		var userID string
		userRows := db.Db.QueryRow(`SELECT user_id from account a where a.google_id=$1`, googleID)
		if userRows.Err() != nil {
			c.AbortWithStatusJSON(500, "User does not exist")
			return
		}
		err = userRows.Scan(&userID)

		db.Db.QueryRow(`DELETE FROM checked_items USING supply_item WHERE user_uuid=$1 AND list_id=$2;`, userID, listID)
		for _, itemID := range checked {
			db.Db.QueryRow(`INSERT INTO checked_items (item_id, user_uuid)
									SELECT si.id, $1 FROM supply_item si
									INNER JOIN supply_list sl on sl.list_id = si.list_id
									WHERE id = $2 AND si.list_id=$3 ON CONFLICT ON CONSTRAINT checked_items_pk DO NOTHING`,
									userID, itemID, listID)
		}

		c.JSON(200, checked)
	}
}

func DeleteSupplyList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		row := db.Db.QueryRow(`DELETE FROM supply_list where list_id=$1`, id)
		if row.Err() != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, nil)
	}
}
