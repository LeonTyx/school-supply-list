package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log"
)

func CheckDBErr(err *pq.Error, c *gin.Context) {
	if err != nil {
		log.Println(err)
		switch err.Code {
		case "23505":
			c.AbortWithStatusJSON(400, "A unique constraint has been violated.")
		case "22004":
			c.AbortWithStatusJSON( 400, "Value cannot be null")
		default:
			c.AbortWithStatusJSON( 503, "There was an error contacting the database.")
		}
		return
	}
}
