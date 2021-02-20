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
		id, err := strconv.Atoi(c.Param("id"))
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
       									CASE WHEN user_uuid IS NOT NULL AND user_uuid = $1
       									    THEN 'true' ELSE 'false' END FROM supply_item sup
										FULL JOIN checked_items ci on sup.id = ci.item_id
										WHERE list_id = $2`, userID, id)
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
