package utils

import (
	"cmpdata/logger"

	"github.com/streadway/amqp"
)

type RMQProducer struct {
	Queue            string
	ConnectionString string
}

func (x RMQProducer) OnError(err error, msg string) {
	if err != nil {
		logger.FatalLn(err.Error())
	}
}

func (x RMQProducer) PublishMessage(contentType string, body []byte) {
	conn, err := amqp.Dial(x.ConnectionString)
	x.OnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	x.OnError(err, "Failed to open a chennel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		x.Queue, //name
		false,   // durable
		false,   // delete when unsued
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	x.OnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // example
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        body,
		})
	x.OnError(err, "Failed to publish a message")
}
