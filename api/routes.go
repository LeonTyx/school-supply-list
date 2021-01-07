package api

import (
	"github.com/gin-gonic/gin"
	"school-supply-list/auth/authorization"
	"school-supply-list/database"
)

//All the routes created by the package nested in
// api/v1/*
func Routes(r *gin.RouterGroup, db *database.DB) {
	//School
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

	//Supply-list
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

	//Supply Item
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
