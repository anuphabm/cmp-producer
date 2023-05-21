package main

import (
	"consumer/config"
	"consumer/database"
	"consumer/logger"
	"consumer/service"
	"testing"
)

func TestService(t *testing.T) {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	database.Init()
	srv := service.NewService("iuat#utbl#^UTBL(\"aaa\",\"bbb\")=1|2|3|4")
	srv.Process()

}
