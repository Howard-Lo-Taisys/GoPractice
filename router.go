package main

import (
	"github.com/gin-gonic/gin"
	. "crawler2/apis"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/crawlerdata", GetDataApi)

	router.POST("/crawlerdata", PostDataApi)

	router.DELETE("/crawlerdata/:id", DeleteDataApi)
   
	return router
}
