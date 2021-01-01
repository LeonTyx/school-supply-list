package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log"
	"school-supply-list/database"
)

type Resource struct {
	ResourceID int    `json:"id"`
	Resource   string `json:"resource"`
	Policy     Policy `json:"policy"`
}

type Policy struct {
	CanAdd    bool `json:"can_add"`
	CanDelete bool `json:"can_delete"`
	CanEdit   bool `json:"can_edit"`
	CanView   bool `json:"can_view"`
}

type Role struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Resources []Resource `json:"resources"`
}

func ValidSession(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := db.SessionStore.Get(c.Request, "session")
		if err != nil {
			c.AbortWithStatusJSON(500, "The server was unable to retrieve this session")
			return
		}

		if session.ID == "" {
			c.AbortWithStatusJSON(401, "This user has no current session. Use of this endpoint is thus unauthorized")
			return
		}
	}
}

func LoadPolicy(db *database.DB, resources string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := db.SessionStore.Get(c.Request, "session")
		if err != nil {
			c.AbortWithStatusJSON(500, "The server was unable to retrieve this session")
			return
		}
		googleID := session.Values["GoogleId"]

		policy, err := getPolicy(db, googleID.(string), resources)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.Set("policy", policy)
	}
}

func getPolicy(db *database.DB, googleID string, resource string) (Policy, error) {
	var policy Policy
	policyQuery, err := db.Db.Query(`SELECT bridge.can_add, bridge.can_view, bridge.can_edit, bridge.can_delete
											FROM account acc INNER JOIN user_role_bridge urb ON acc.user_id = urb.user_uuid
											LEFT JOIN role_resource_bridge bridge ON bridge.role_id = urb.role_id
											INNER JOIN resource rsc ON bridge.resource_id = rsc.resource_id
											WHERE google_id=$1 AND rsc.resource_name=$2`, googleID, resource)
	if err != nil {
		log.Fatal(err)
		return policy, err
	}
	defer policyQuery.Close()
	for policyQuery.Next() {
		var temp Policy
		err = policyQuery.Scan(&temp.CanAdd, &temp.CanView,
			&temp.CanEdit, &temp.CanDelete)

		policy.CanAdd = temp.CanAdd || policy.CanAdd
		policy.CanDelete = temp.CanDelete || policy.CanDelete
		policy.CanEdit = temp.CanEdit || policy.CanEdit
		policy.CanView = temp.CanView || policy.CanView
	}

	return policy, nil
}

func CanView() gin.HandlerFunc {
	return func(c *gin.Context) {
		policyStr, exists := c.Get("policy")
		if !exists {
			c.AbortWithStatusJSON(500, "Policy was not loaded.")
		}
		policy := policyStr.(Policy)
		if !policy.CanView {
			c.AbortWithStatusJSON(403, "You do not have access to this endpoint.")
			return
		}
	}
}

func CanCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		policyStr, exists := c.Get("policy")
		if !exists {
			c.AbortWithStatusJSON(500, "Policy was not loaded.")
		}
		policy := policyStr.(Policy)

		if !policy.CanAdd {
			c.AbortWithStatusJSON(403, "You do not have access to this endpoint.")
			return
		}
	}
}

func CanEdit() gin.HandlerFunc {
	return func(c *gin.Context) {
		policyStr, exists := c.Get("policy")
		if !exists {
			c.AbortWithStatusJSON(500, "Policy was not loaded.")
		}
		policy := policyStr.(Policy)

		if !policy.CanEdit {
			c.AbortWithStatusJSON(403, "You do not have access to this endpoint.")
			return
		}
	}
}

func CanDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		policyStr, exists := c.Get("policy")
		if !exists {
			c.AbortWithStatusJSON(500, "Policy was not loaded.")
		}
		policy := policyStr.(Policy)

		if !policy.CanDelete {
			c.AbortWithStatusJSON(403, "You do not have access to this endpoint.")
			return
		}
	}
}
