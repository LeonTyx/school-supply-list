package events

import (
	"github.com/gin-gonic/gin"
	"school-supply-list/database"
	"time"
)

type event struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Dates []date `json:"dates"`
}

type date struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}


func CreateEvent(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var event event
		err := c.BindJSON(&event)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}

		row := db.Db.QueryRow(`INSERT INTO school (school_name) VALUES ($1) returning school_id`, event.Title, event.Desc)
		err = row.Scan(&event.ID)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid event title and description.")
			return
		}

		c.JSON(200, event)
	}
}

func GetEvent(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var event event

		c.JSON(200, event)
	}
}

func GetEvents(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var events []event

		c.JSON(200, events)
	}
}

func UpdateEvent(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var event event
		err := c.BindJSON(&event)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}

		c.JSON(200, event)
	}
}

func DeleteEvent(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var event event
		err := c.BindJSON(&event)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}

		c.JSON(200, event)
	}
}