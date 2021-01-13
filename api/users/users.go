package users

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/auth/authorization"
	"school-supply-list/database"
	"strconv"
)

func GetUser(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		role := authorization.Role{
			ID: id,
		}

		row := db.Db.QueryRow(`SELECT role_id, role_name, role_desc FROM role
											WHERE role.role_id=$1`, id)
		err = row.Scan(&role.ID, &role.Name, &role.Desc)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, role)
	}
}

func GetAllUsers(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roles []authorization.Role
		rows, err := db.Db.Query(`SELECT role_id, role_name, role_desc FROM role`)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		for rows.Next() {
			var role authorization.Role
			err = rows.Scan(&role.ID, &role.Name, &role.Desc)
			roles = append(roles, role)
		}
		c.JSON(200, roles)
	}
}

func UpdateUser(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		var role authorization.Role
		err = c.BindJSON(&role)
		role.ID = id
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}

		row := db.Db.QueryRow(`UPDATE role SET role_name=$1, role_desc=$2
	   		where role_id=$3 returning role_id, role_name, role_desc`, role.Name, role.Desc)
		//Scan the latest changes into the role struct
		err = row.Scan(&role.ID, &role.Name, role.Desc)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		for _, resourceDetails := range role.Resources{
			row = db.Db.QueryRow(`UPDATE role_resource_bridge SET can_add=$1, can_view=$2, 
                                can_edit=$3, can_delete=$4, resource_id=$5 WHERE role_id=$6`,
				resourceDetails.Policy.CanAdd, resourceDetails.Policy.CanView,
				resourceDetails.Policy.CanEdit, resourceDetails.Policy.CanDelete,
				resourceDetails.ResourceID, role.ID)
			if row.Err() != nil {
				database.CheckDBErr(err.(*pq.Error), c)
				return
			}
		}
		c.JSON(200, role)
	}
}

func DeleteUser(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		row := db.Db.QueryRow(`DELETE FROM role where role_id=$1 `, id)
		if row.Err() != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, nil)
	}
}

