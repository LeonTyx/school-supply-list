package api

import (
	"github.com/gin-gonic/gin"
	"school-supply-list/database"
)

//All the routes created by the package nested in
// oauth/v1/*
func Routes(r *gin.RouterGroup, db *database.DB) {
	r.PUT("/school", createSchool(db))
	r.GET("/school", getSchool(db))
	r.POST("/school", updateSchool(db))
	r.DELETE("/school", deleteSchool(db))
}

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
