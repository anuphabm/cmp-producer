package main

import (
	"cmpdata/config"
	"cmpdata/logger"
	"cmpdata/router"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Logger Initialized successfully")
	router.Init()
	logger.InfoLn("Router Initialized successfully")
}
