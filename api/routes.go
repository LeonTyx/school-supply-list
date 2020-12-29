package api

import (
	"github.com/gin-gonic/gin"
	"school-supply-list/auth/authorization"
	"school-supply-list/database"
)

//All the routes created by the package nested in
// api/v1/*
func Routes(r *gin.RouterGroup, db *database.DB) {
	r.PUT("/school", authorization.ValidSession(db),
		authorization.CanCreate(),
		createSchool(db))
	r.GET("/school",
		getSchool(db))
	r.POST("/school", authorization.ValidSession(db),
		authorization.CanEdit(),
		updateSchool(db))
	r.DELETE("/school", authorization.ValidSession(db),
		authorization.CanDelete(),
		deleteSchool(db))

	r.PUT("/supply-list", authorization.ValidSession(db),
		authorization.CanCreate(),
		createSupplyList(db))
	r.GET("/supply-list", getSupplyList(db))
	r.POST("/supply-list", authorization.ValidSession(db),
		authorization.CanEdit(),
		updateSupplyList(db))
	r.DELETE("/supply-list", authorization.ValidSession(db),
		authorization.CanDelete(),
		deleteSupplyList(db))
}
