package main

import (
	"consumer/config"
	"consumer/consts"
	"consumer/database"
	"consumer/handlers"
	"consumer/logger"
	"consumer/utils"
	"fmt"
	"os"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Logger Initialized successfully")
	database.Init()
	logger.InfoLn("Database Initialize successfully")

	// connection to rmq
	envMode := os.Getenv("RUN_MODE")
	connectionString := config.Appconfig.GetString(fmt.Sprintf("%s.server.rmqurl", envMode))
	exampleQueue := utils.RMQConsumer{
		Queue:            consts.EXAMPLE_QUEUE,
		ConnectionString: connectionString,
		MsgHandler:       handlers.HandleExample,
	}
	// Start consuming message on the specified queues
	forever := make(chan bool)

	go exampleQueue.Consume()

	// Multiple listeners can be specified here

	<-forever

}
