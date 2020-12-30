package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/database"
	"strconv"
)

type school struct {
	SchoolID int `json:"school_id"`
	SchoolName string `json:"school_name"`
}

func createSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(200, gin.H{
			"id": id,
			"body": c.PostForm("school_name"),
		})
	}
}

func getSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer")
		}
		school := school{
			SchoolID: id,
		}

		schoolRows, err := db.Db.Query(`SELECT school_name from school 
											where school.school_id=$1`, id)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
		}

		for schoolRows.Next() {
			err = schoolRows.Scan(&school.SchoolName)
			if err != nil {
				c.AbortWithStatusJSON(500, "The server was unable to retrieve school info")
			}
		}

		c.JSON(200, school)
	}
}

func getSchools(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var schools []school
		schoolsRows, err := db.Db.Query(`SELECT school_id, school_name from school`)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
		}

		for schoolsRows.Next() {
			var school school
			err = schoolsRows.Scan(&school.SchoolID, &school.SchoolName)
			if err != nil {
				c.AbortWithStatusJSON(500, "The server was unable to retrieve school info")
			}
			schools = append(schools, school)
		}

		c.JSON(200, schools)
	}
}

func updateSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, c.Request.Body)
	}
}

func deleteSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer")
		}

		_, err = db.Db.Query(`DELETE from school where school.school_id=$1`, id)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
		}

		c.JSON(200, nil)
	}
}
