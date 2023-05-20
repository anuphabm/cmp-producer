package router

import (
	"cmpdata/api/controller"
	"cmpdata/config"
	"cmpdata/middleware"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()
	router.Run(config.Appconfig.GetString("server.port"))
}

func NewRouter() *gin.Engine {
	router := gin.New()
	resource := router.Group("/api")
	resource.Use(middleware.LogRequestInfo())
	{
		resource.POST("/env/:env/table/:table", controller.ReciveData)
	}
	return router
}
