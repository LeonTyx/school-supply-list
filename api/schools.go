package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/database"
	"strconv"
)

type school struct {
	SchoolID   int    `json:"school_id"`
	SchoolName string `json:"school_name"`
}

func createSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var school school
		err := c.BindJSON(&school)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}
		row := db.Db.QueryRow(`INSERT INTO school (school_name) VALUES ($1) returning school_id`, school.SchoolName)
		err = row.Scan(&school.SchoolID)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid school name.")
			return
		}

		c.JSON(200, school)
	}
}

func getSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer")
			return
		}
		school := school{
			SchoolID: id,
		}

		schoolRows, err := db.Db.Query(`SELECT school_name from school 
											where school.school_id=$1`, id)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
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
			return
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
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer")
			return
		}

		var school school
		err = c.BindJSON(&school)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}

		row := db.Db.QueryRow(`UPDATE school SET school_name = $1 WHERE school_id=$2`, school.SchoolName, id)
		if row.Err() != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		c.JSON(200, school.SchoolID)
	}
}

func deleteSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer")
			return
		}

		row := db.Db.QueryRow(`DELETE from school where school.school_id=$1`, id)
		if row.Err() != nil {
			database.CheckDBErr(err.(*pq.Error), c)
		}

		c.JSON(200, nil)
	}
}
