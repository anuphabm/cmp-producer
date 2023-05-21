package handlers

import (
	"consumer/logger"
	"consumer/service"
	"fmt"

	"github.com/streadway/amqp"
)

func HandleExample(queue string, msg amqp.Delivery, err error) {
	if err != nil {
		msgErr := fmt.Sprintf("Error occurred in RMQ consumer : %v", err)
		logger.FatalLn(msgErr)
	}
	msgInfo := fmt.Sprintf("Message received on '%s' queue: %s", queue, string(msg.Body))
	logger.InfoLn(msgInfo)

	srv := service.NewService(string(msg.Body))
	srv.Process()
}
