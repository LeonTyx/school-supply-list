package users

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/database"
)

type User struct {
	ID            string   `json:"user_id"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	AccountImgURL string   `json:"account_img_url"`
	Roles         []int `json:"roles"`
}

func GetUser(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user := User{
			ID: id,
		}

		row := db.Db.QueryRow(`SELECT name, email, google_picture FROM account
											WHERE user_id=$1`, id)
		err := row.Scan(&user.Name, &user.Email, &user.AccountImgURL)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, user)
	}
}

func GetAllUsers(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []User
		rows, err := db.Db.Query(`SELECT name, email, google_picture FROM account`)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		for rows.Next() {
			var user User
			err = rows.Scan(&user.Name, &user.Email, &user.AccountImgURL)
			if err != nil {
				database.CheckDBErr(err.(*pq.Error), c)
				return
			}
			users = append(users, user)
		}
		c.JSON(200, users)
	}
}

func UpdateUser(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var user User
		err := c.BindJSON(&user)
		user.ID = id
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}
		row := db.Db.QueryRow(`DELETE FROM user_role_bridge WHERE user_uuid=$1`, id)
		if row.Err() != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		for role := range user.Roles{
			row = db.Db.QueryRow(`INSERT INTO user_role_bridge (user_uuid, role_id) VALUES ($1, $2)`, user.ID, role)
			if row.Err() != nil {
				database.CheckDBErr(err.(*pq.Error), c)
				return
			}
		}
		c.JSON(200, user)
	}
}

func DeleteUser(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		row := db.Db.QueryRow(`DELETE FROM account where user_id=$1 `, id)
		if row.Err() != nil {
			database.CheckDBErr(row.Err().(*pq.Error), c)
			return
		}

		c.JSON(200, nil)
	}
}
