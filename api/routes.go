package api

import (
	"github.com/gin-gonic/gin"
	"school-supply-list/database"
)

//All the routes created by the package nested in
// oauth/v1/*
func Routes(r *gin.RouterGroup, db *database.DB) {
	r.PUT("/school", createSchool(db))
	r.GET("/school", getSchool(db))
	r.POST("/school", updateSchool(db))
	r.DELETE("/school", deleteSchool(db))

	r.PUT("/supply-list", createSupplyList(db))
	r.GET("/supply-list", getSupplyList(db))
	r.POST("/supply-list", updateSupplyList(db))
	r.DELETE("/supply-list", deleteSupplyList(db))
}