package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/database"
)

type ResourceScope struct {
	Resource map[string]string
	Scope    map[string]string
	ECAID    string
}

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

func ResourceCtx(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		resourceScope := c.Request.Context().Value("resource_scope")

		session, err := db.SessionStore.Get(c.Request, "session")
		if err != nil {
			c.AbortWithStatusJSON(500, "The server was unable to retrieve this session")
			return
		}
		googleID := session.Values["GoogleId"]

		policy, err := getPolicy(googleID.(string), resourceScope.(ResourceScope))
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.Set("policy", policy)
	}
}

func getPolicy(googleID string, resourceScope ResourceScope) (Policy, error) {
	var memberPolicy Policy

	policy := Policy{
		CanAdd:    memberPolicy.CanAdd,
		CanView:   memberPolicy.CanView,
		CanEdit:   memberPolicy.CanEdit,
		CanDelete: memberPolicy.CanDelete,
	}

	return policy, nil
}

func CanView() gin.HandlerFunc {
	return func(c *gin.Context) {
		policy := c.Request.Context().Value("policy").(Policy)

		if !policy.CanView {
			c.AbortWithStatusJSON(403, "You do not have access to this endpoint.")
			return
		}
	}
}

func CanCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		policy := c.Request.Context().Value("policy").(Policy)

		if !policy.CanAdd {
			c.AbortWithStatusJSON(403, "You do not have access to this endpoint.")
			return
		}
	}
}

func CanEdit() gin.HandlerFunc {
	return func(c *gin.Context) {
		policy := c.Request.Context().Value("policy").(Policy)

		if !policy.CanEdit {
			c.AbortWithStatusJSON(403, "You do not have access to this endpoint.")
			return
		}
	}
}

func CanDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		policy := c.Request.Context().Value("policy").(Policy)

		if !policy.CanDelete {
			c.AbortWithStatusJSON(403, "You do not have access to this endpoint.")
			return
		}
	}
}
