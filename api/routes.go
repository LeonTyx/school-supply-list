package api

import (
	"github.com/gin-gonic/gin"
	"school-supply-list/api/permissions"
	"school-supply-list/api/schools"
	"school-supply-list/api/supplies"
	"school-supply-list/api/supplylist"
	"school-supply-list/api/users"
	"school-supply-list/auth/authorization"
	"school-supply-list/database"
)

//All the routes created by the package nested in
// api/v1/*
func Routes(r *gin.RouterGroup, db *database.DB) {
	schoolRoutes(r, db)
	listRoutes(r, db)
	supplyRoutes(r, db)
	roleRoutes(r, db)
	userRoutes(r, db)
	resourceRoute(r, db)
}

func resourceRoute(r *gin.RouterGroup, db *database.DB) {
	r.GET("/resource",
		authorization.LoadPolicy(db, "resources"),
		authorization.CanView(),
		permissions.GetAllResources(db))
}

func schoolRoutes(r *gin.RouterGroup, db *database.DB) {
	r.PUT("/school", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "school"),
		authorization.CanCreate(),
		schools.CreateSchool(db))
	r.GET("/school/:id",
		schools.GetSchool(db))
	r.GET("/schools",
		schools.GetSchools(db))
	r.POST("/school/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "school"),
		authorization.CanEdit(),
		schools.UpdateSchool(db))
	r.DELETE("/school/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "school"),
		authorization.CanDelete(),
		schools.DeleteSchool(db))
}

func listRoutes(r *gin.RouterGroup, db *database.DB) {
	r.PUT("/supply-list", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply-list"),
		authorization.CanCreate(),
		supplylist.CreateSupplyList(db))
	r.GET("/supply-list/:id",
		supplylist.GetSupplyList(db))
	r.GET("/supply-lists",
		supplylist.GetSupplyLists(db))
	r.POST("/supply-list/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply-list"),
		authorization.CanEdit(),
		supplylist.UpdateSupplyList(db))
	r.DELETE("/supply-list/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply-list"),
		authorization.CanDelete(),
		supplylist.DeleteSupplyList(db))
}

func supplyRoutes(r *gin.RouterGroup, db *database.DB) {
	r.PUT("/supply", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply"),
		authorization.CanCreate(),
		supplies.CreateSupply(db))
	r.GET("/supply/:id",
		authorization.LoadPolicy(db, "supply"),
		authorization.CanView(),
		supplies.GetSupply(db))
	r.GET("/supply",
		authorization.LoadPolicy(db, "supply"),
		authorization.CanView(),
		supplies.GetAllSupplies(db))
	r.POST("/supply/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply"),
		authorization.CanEdit(),
		supplies.UpdateSupply(db))
	r.DELETE("/supply/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply"),
		authorization.CanDelete(),
		supplies.DeleteSupply(db))
}

func roleRoutes(r *gin.RouterGroup, db *database.DB) {
	r.PUT("/role", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "role"),
		authorization.CanCreate(),
		permissions.CreateRole(db))
	r.GET("/role/:id",
		authorization.LoadPolicy(db, "role"),
		authorization.CanView(),
		permissions.GetRole(db))
	r.GET("/role",
		authorization.LoadPolicy(db, "role"),
		authorization.CanView(),
		permissions.GetAllRoles(db))
	r.POST("/role/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "role"),
		authorization.CanEdit(),
		permissions.UpdateRole(db))
	r.DELETE("/role/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "role"),
		authorization.CanDelete(),
		permissions.DeleteRole(db))
}

func userRoutes(r *gin.RouterGroup, db *database.DB) {
	r.GET("/user/:id",
		authorization.LoadPolicy(db, "user"),
		authorization.CanView(),
		users.GetUser(db))
	r.GET("/user",
		authorization.LoadPolicy(db, "user"),
		authorization.CanView(),
		users.GetAllUsers(db))
	r.POST("/user/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "user"),
		authorization.CanEdit(),
		users.UpdateUser(db))
	r.DELETE("/user/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "user"),
		authorization.CanDelete(),
		users.DeleteUser(db))
}