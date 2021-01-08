package api

import (
	"github.com/gin-gonic/gin"
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
}

func schoolRoutes(r *gin.RouterGroup, db *database.DB) {
	r.PUT("/school", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "school"),
		authorization.CanCreate(),
		createSchool(db))
	r.GET("/school/:id",
		getSchool(db))
	r.GET("/schools",
		getSchools(db))
	r.POST("/school/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "school"),
		authorization.CanEdit(),
		updateSchool(db))
	r.DELETE("/school/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "school"),
		authorization.CanDelete(),
		deleteSchool(db))
}

func listRoutes(r *gin.RouterGroup, db *database.DB) {
	r.PUT("/supply-list", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply-list"),
		authorization.CanCreate(),
		createSupplyList(db))
	r.GET("/supply-list/:id",
		getSupplyList(db))
	r.GET("/supply-lists",
		getSupplyLists(db))
	r.POST("/supply-list/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply-list"),
		authorization.CanEdit(),
		updateSupplyList(db))
	r.DELETE("/supply-list/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply-list"),
		authorization.CanDelete(),
		deleteSupplyList(db))
}

func supplyRoutes(r *gin.RouterGroup, db *database.DB) {
	r.PUT("/supply", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply"),
		authorization.CanCreate(),
		createSupply(db))
	r.GET("/supply/:id",
		authorization.LoadPolicy(db, "supply"),
		authorization.CanView(),
		getSupply(db))
	r.GET("/supply",
		authorization.LoadPolicy(db, "supply"),
		authorization.CanView(),
		getAllSupplies(db))
	r.POST("/supply/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply"),
		authorization.CanEdit(),
		updateSupply(db))
	r.DELETE("/supply/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply"),
		authorization.CanDelete(),
		deleteSupply(db))
}

func roleRoutes(r *gin.RouterGroup, db *database.DB) {
	r.PUT("/role", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply"),
		authorization.CanCreate(),
		createRole(db))
	r.GET("/role/:id",
		authorization.LoadPolicy(db, "supply"),
		authorization.CanView(),
		getRole(db))
	r.GET("/role",
		authorization.LoadPolicy(db, "supply"),
		authorization.CanView(),
		getAllRoles(db))
	r.POST("/role/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply"),
		authorization.CanEdit(),
		updateRole(db))
	r.DELETE("/role/:id", authorization.ValidSession(db),
		authorization.LoadPolicy(db, "supply"),
		authorization.CanDelete(),
		deleteRole(db))
}