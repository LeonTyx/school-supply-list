package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/database"
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
		}

		c.JSON(200, list)
	}
}

func getSupplyList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": id,
		})
	}
}

func getSupplyLists(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "we did it!",
		})
	}
}

func updateSupplyList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "we did it!",
		})
	}
}

func deleteSupplyList(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "we did it!",
		})
	}
}
