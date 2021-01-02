package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/database"
	"strconv"
)

type supplyList struct {
	ListID    int        `json:"list-id"`
	Grade     int        `json:"grade"`
	SchoolID  int        `json:"school_id"`
	ListName  string     `json:"list-name"`
	ListItems []listItem `json:"list-items"`
}

type listItem struct {
	ItemID   int    `json:"item-id"`
	ItemName string `json:"item-name"`
	ItemDesc string `json:"item-desc"`
}

func createSupplyList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var list supplyList
		err := json.NewDecoder(c.Request.Body).Decode(&list)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}
		row := db.Db.QueryRow(`INSERT INTO supply_list (grade, list_name, school_id, list_id) 
		  VALUES ($1, $2, $3, default) returning list_id`, list.Grade, list.ListName, list.ListID)
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

		row := db.Db.QueryRow(`SELECT grade, list_name, school_id from supply_list 
											where supply_list.list_id=$1`, id)
		err = row.Scan(&list.Grade, &list.ListName, &list.SchoolID)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, list)
	}
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
		row := db.Db.QueryRow(`UPDATE supply_list SET grade=$1, list_name=$2, 
	   		school_id=$3 where list_id=$4 returning list_id`, list.Grade, list.ListName, list.SchoolID, list.ListID)
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
		if row.Err() != nil{
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, nil)
	}
}
