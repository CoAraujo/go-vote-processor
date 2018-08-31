package stream

import (
	"fmt"
	"log"

	config "github.com/coaraujo/go-vote-processor/config/rabbit"
	"github.com/streadway/amqp"
)

type RabbitStream struct {
	RabbitMQ   *config.RabbitMQ
	Queue      amqp.Queue
	QueueGroup string
}

func NewRabbitStream(conn *config.RabbitMQ, queue string) *RabbitStream {
	q := "vote.groupBBB"
	if queue != "" {
		q = queue
	}
	rabbitStream := RabbitStream{RabbitMQ: conn, QueueGroup: q}
	rabbitStream.Queue = conn.CreateQueue(q)
	return &rabbitStream
}

func (r *RabbitStream) ReceiveVote() {
	fmt.Println("[RABBITMQ] Receiving vote...")

	votes, err := r.RabbitMQ.GetChannel().Consume(
		r.Queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for vote := range votes {
			log.Printf("Received a message: %s", vote.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
