package api

import (
	"github.com/gin-gonic/gin"
	"school-supply-list/database"
)

func createSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "we did it!",
		})
	}
}

func getSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "we did it!",
		})
	}
}

func updateSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "we did it!",
		})
	}
}

func deleteSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "we did it!",
		})
	}
}
