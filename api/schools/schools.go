package schools

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

func CreateSchool(db *database.DB) gin.HandlerFunc {
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

type schoolWithList struct {
	SchoolID   int           `json:"school_id"`
	SchoolName string        `json:"school_name"`
	SupplyList []listDetails `json:"supply_lists"`
}

type listDetails struct {
	ListID   int    `json:"list_id"`
	Grade    int    `json:"grade"`
	SchoolID int    `json:"school_id"`
	ListName string `json:"list_name"`
}

func GetSchool(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer")
			return
		}

		var school schoolWithList
		schoolRows, err := db.Db.Query(`SELECT school_name, school_id from school 
											where school.school_id=$1`, id)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		rowcount := 0
		for schoolRows.Next() {
			err = schoolRows.Scan(&school.SchoolName, &school.SchoolID)
			if err != nil {
				c.AbortWithStatusJSON(500, "The server was unable to retrieve school info")
				return
			}
			rowcount++
		}
		if rowcount == 0 {
			c.AbortWithStatusJSON(404, "This school does not exist")
			return
		}

		listRows, err := db.Db.Query(`SELECT list_id, grade, list_name, school_id from supply_list`)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		for listRows.Next() {
			var list listDetails
			err = listRows.Scan(&list.ListID, &list.Grade, &list.ListName, &list.SchoolID)
			school.SupplyList = append(school.SupplyList, list)
		}

		c.JSON(200, school)
	}
}

func GetSchools(db *database.DB) gin.HandlerFunc {
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

func UpdateSchool(db *database.DB) gin.HandlerFunc {
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

func DeleteSchool(db *database.DB) gin.HandlerFunc {
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
