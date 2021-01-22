package permissions

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/auth/authorization"
	"school-supply-list/database"
	"strconv"
)

func GetAllResources(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		resources := make(map[string]authorization.Resource)
		rows, err := db.Db.Query(`SELECT resource_id, resource_name FROM resource`)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		for rows.Next() {
			var resource authorization.Resource
			var resourceName string
			err = rows.Scan(&resource.ResourceID, &resourceName)
			resources[resourceName] = resource
		}

		c.JSON(200, resources)
	}
}

func CreateRole(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var role authorization.Role
		err := json.NewDecoder(c.Request.Body).Decode(&role)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}
		row := db.Db.QueryRow(`INSERT INTO role (role_name, role_desc) 
		  VALUES ($1, $2) RETURNING role_id`, role.Name, role.Desc)
		err = row.Scan(&role.ID)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		for _, resourceDetails := range role.Resources {
			row = db.Db.QueryRow(`INSERT INTO role_resource_bridge (can_add, can_view, can_edit, can_delete, resource_id, role_id) 
		  			VALUES ($1, $2, $3, $4, $5, $6)`, resourceDetails.Policy.CanAdd, resourceDetails.Policy.CanView,
				resourceDetails.Policy.CanEdit, resourceDetails.Policy.CanDelete, resourceDetails.ResourceID, role.ID)
			if row.Err() != nil {
				database.CheckDBErr(err.(*pq.Error), c)
				return
			}
		}

		c.JSON(200, role)
	}
}

func GetRole(db *database.DB) gin.HandlerFunc {
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

		err, role.Resources = getPermissions(role.ID, db)
		if err != nil{
			c.AbortWithStatusJSON(500, "Unable to retrieve permissions for this role")
		}

		c.JSON(200, role)
	}
}

func GetAllRoles(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roles = make(map[int]authorization.Role)
		rows, err := db.Db.Query(`SELECT role_id, role_name, role_desc FROM role`)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		for rows.Next() {
			var role authorization.Role
			err = rows.Scan(&role.ID, &role.Name, &role.Desc)
			if err != nil {
				database.CheckDBErr(err.(*pq.Error), c)
				return
			}
			err, role.Resources = getPermissions(role.ID, db)
			if err != nil{
				c.AbortWithStatusJSON(500, "Unable to retrieve permissions for this role")
			}
			roles[role.ID] = role;
		}
		c.JSON(200, roles)
	}
}

func getPermissions(roleID int, db *database.DB) (error, map[string]authorization.Resource){
	resourceMap := make(map[string]authorization.Resource)
	resourceRows, err := db.Db.Query(`SELECT resource.resource_id, resource.resource_name, can_add, can_delete, can_view, can_edit FROM resource 
    								JOIN role_resource_bridge rrb on resource.resource_id = rrb.resource_id
    								WHERE rrb.role_id=$1`, roleID)
	if err != nil {
		return err.(*pq.Error), resourceMap
	}

	for resourceRows.Next(){
		var resource authorization.Resource
		var resourceName string
		err = resourceRows.Scan(&resource.ResourceID, &resourceName, &resource.Policy.CanAdd,
			&resource.Policy.CanDelete, &resource.Policy.CanView, &resource.Policy.CanEdit)

		if err != nil {
			return err.(*pq.Error), resourceMap
		}
		resourceMap[resourceName] = resource
	}

	return nil, resourceMap
}

func UpdateRole(db *database.DB) gin.HandlerFunc {
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
	   		where role_id=$3 returning role_id, role_name, role_desc`, role.Name, role.Desc, role.ID)
		//Scan the latest changes into the role struct
		err = row.Scan(&role.ID, &role.Name, &role.Desc)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		for _, resourceDetails := range role.Resources {
			row = db.Db.QueryRow(`INSERT INTO role_resource_bridge (can_add, can_view, can_edit, can_delete, resource_id, role_id) 
																		VALUES ($1, $2, $3, $4, $5, $6)
										ON CONFLICT (resource_id, role_id) 
										    DO UPDATE SET can_add=$1, can_view=$2, can_edit=$3, can_delete=$4 
										WHERE role_resource_bridge.resource_id=$5 AND role_resource_bridge.role_id=$6`,
				resourceDetails.Policy.CanAdd, resourceDetails.Policy.CanView,
				resourceDetails.Policy.CanEdit, resourceDetails.Policy.CanDelete,
				resourceDetails.ResourceID, role.ID)
			if row.Err() != nil {
				database.CheckDBErr(row.Err().(*pq.Error), c)
				return
			}
		}
		c.JSON(200, role)
	}
}

func DeleteRole(db *database.DB) gin.HandlerFunc {
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
