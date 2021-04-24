package comms

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type MQClient struct {
	conn *amqp.Connection
}

func CreateMQClient() MQClient{
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}
	return MQClient{conn}
}

func (mq MQClient) SubscribeToQueue(queueName string, channel chan []byte) {
	ch, _ := mq.conn.Channel()

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
			channel <- d.Body
		}
	}()
	<-forever
}

