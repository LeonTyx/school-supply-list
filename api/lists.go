package api

import (
	"github.com/gin-gonic/gin"
	"school-supply-list/database"
)

type supplyList struct {
	ListID    int        `json:"list-id"`
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
		c.JSON(200, gin.H{
			"message": "we did it!",
		})
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
