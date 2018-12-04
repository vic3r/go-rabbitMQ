package main

import (
	"log"

	amqp "github.com/streadway/amqp"
)

func connect() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	queue := initQueue(ch)
	publish(ch, queue)

}

func initQueue(ch *amqp.Channel) *amqp.Queue {
	queue, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}
	return &queue
}

func publish(ch *amqp.Channel, queue *amqp.Queue) {
	body := "Toledo se la come!"
	err := ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message %v", err)
	}
}

func main() {
	connect()
}
