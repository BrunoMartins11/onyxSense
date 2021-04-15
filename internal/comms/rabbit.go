package comms

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var channel *amqp.Channel

func CreateRabbitChannel() {
	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("RABBIT_SERVER_URL")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channel, err = connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
}

func SubscribeQueue	(){

	// Subscribing to QueueService1 for getting messages.
	messages, err := channel.Consume(
		"QueueService1", // queue name
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message.Body)
		}
	}()

	<-forever
}
