package main

import (
	"log"

	"github.com/streadway/amqp"
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
	consume(ch, queue)
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
		log.Fatalf("Failed to declare queue")
	}

	return &queue
}

func consume(ch *amqp.Channel, queue *amqp.Queue) {
	msgs, err := ch.Consume(
		queue.Name, // queueu
		"",         // consumer
		true,       // auto-ack
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to register consumer %v", err)
	}

	listen := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-listen
}

func main() {
	connect()
}
