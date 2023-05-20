package main

import (
	"cmpdata/config"
	"cmpdata/database"
	"cmpdata/logger"
	"cmpdata/router"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Logger Initialized successfully")
	database.Init()
	logger.InfoLn("Database Initialize successfully")
	router.Init()
	logger.InfoLn("Router Initialized successfully")
}
