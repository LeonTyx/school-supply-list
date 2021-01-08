package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/database"
	"strconv"
)

type supplyList struct {
	ListID              int                     `json:"list_id"`
	Grade               int                     `json:"grade"`
	SchoolID            int                     `json:"school_id"`
	ListName            string                  `json:"list_name"`
	BasicSupplies       []supplyItem            `json:"basic_supplies"`
	CategorizedSupplies map[string][]supplyItem `json:"categorized_supplies"`
	Published           bool                    `json:"published"`
}

func createSupplyList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var list supplyList
		err := json.NewDecoder(c.Request.Body).Decode(&list)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}
		row := db.Db.QueryRow(`INSERT INTO supply_list (grade, list_name, school_id, published, list_id) 
		  VALUES ($1, $2, $3, $4, default) returning list_id`, list.Grade, list.ListName, list.ListID, list.Published)
		err = row.Scan(&list.ListID)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, list)
	}
}

func getSupplyList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		list := supplyList{
			ListID: id,
		}

		row := db.Db.QueryRow(`SELECT grade, list_name, published, school_id from supply_list 
											where supply_list.list_id=$1`, id)
		err = row.Scan(&list.Grade, &list.ListName, &list.SchoolID)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		list.BasicSupplies, list.CategorizedSupplies,err = getItemsForList(list.ListID, db)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, list)
	}
}

func getItemsForList(id int, db *database.DB) ([]supplyItem, map[string][]supplyItem, error) {
	var basicSupplies []supplyItem
	categorizedSupplies := make(map[string][]supplyItem)
	rows, err := db.Db.Query(`SELECT id, supply_name, supply_desc, ilb.category FROM supply_item sup 
										INNER JOIN item_list_bridge ilb on sup.id = ilb.item_id
										WHERE ilb.list_id = $1`, id)
	if err != nil {
		return basicSupplies, categorizedSupplies, err
	}
	for rows.Next() {
		var supply supplyItem

		err = rows.Scan(&supply.ID, &supply.Supply, &supply.Desc, &supply.Category)
		basicSupplies = append(basicSupplies, supply)

		// Check if item is categorized
		if supply.Category.Valid {
			// Either create a new map item or add to existing item
			if val, ok := categorizedSupplies[supply.Category.String]; ok {
				val = append(val, supply)
			}else{
				categorizedSupplies[supply.Category.String] = []supplyItem{supply}
			}
		}else{
			basicSupplies = append(basicSupplies, supply)
		}
	}
	return basicSupplies, categorizedSupplies, nil
}

func getSupplyLists(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var lists []supplyList
		rows, err := db.Db.Query(`SELECT grade, list_name, school_id from supply_list`)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		for rows.Next() {
			var list supplyList
			err = rows.Scan(&list.Grade, &list.ListName, &list.SchoolID)
			lists = append(lists, list)
		}
		c.JSON(200, lists)
	}
}

func updateSupplyList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		var list supplyList
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

func deleteSupplyList(db *database.DB) gin.HandlerFunc {
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
