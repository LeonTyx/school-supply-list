package api

import (
	"github.com/gin-gonic/gin"
	"school-supply-list/database"
)

func Test(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "we did it!",
		})
	}
}
