package router

import (
	"cmpdata/api/controller"
	"cmpdata/config"
	"cmpdata/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()
	envMode := os.Getenv("RUN_MODE")
	router.Run(config.Appconfig.GetString(fmt.Sprintf("%s.server.port", envMode)))
}

func NewRouter() *gin.Engine {
	router := gin.New()
	envMode := os.Getenv("RUN_MODE")
	group := config.Appconfig.GetString(fmt.Sprintf("%s.server.path", envMode))
	resource := router.Group(group)
	resource.Use(middleware.LogRequestInfo())
	{
		resource.POST("/env/:env/table/:table", controller.ReciveData)
	}
	return router
}
