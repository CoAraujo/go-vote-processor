package stream

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	config "github.com/coaraujo/go-vote-processor/config/rabbit"
	domain "github.com/coaraujo/go-vote-processor/domain"
	"github.com/streadway/amqp"
)

const (
	queueGroup = "vote.groupBBB"
)

type RabbitStream struct {
	RabbitMQ *config.RabbitMQ
	Queue    amqp.Queue
}

func NewRabbitStream(conn *config.RabbitMQ) *RabbitStream {
	rabbitStream := RabbitStream{RabbitMQ: conn}
	rabbitStream.Queue = conn.CreateQueue(queueGroup)
	return &rabbitStream
}

func (r *RabbitStream) SendVote(vote domain.Vote) {
	fmt.Println("[RABBITMQ] Sending vote...")

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(vote)
	reqBodyBytes.Bytes()

	err := r.RabbitMQ.GetChannel().Publish(
		"",         // exchange
		queueGroup, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(reqBodyBytes.Bytes()),
		})
	failOnError(err, "[RABBITMQ] Failed to publish a message")

	fmt.Println("[RABBITMQ] Vote sent sucessfully.")
}

//TESTAR O RECEIVE VOTE
// func (r *RabbitStream) ReceiveVote(vote domain.Vote) {
// 	fmt.Println("[RABBITMQ] Receiving vote...")

// 	votes, err := r.RabbitMQ.GetChannel().Consume(
// 		r.Queue.Name, // queue
// 		"",           // consumer
// 		true,         // auto-ack
// 		false,        // exclusive
// 		false,        // no-local
// 		false,        // no-wait
// 		nil,          // args
// 	)
// 	failOnError(err, "Failed to register a consumer")

// 	forever := make(chan bool)

// 	go func() {
// 		for vote := range votes {
// 			log.Printf("Received a message: %s", vote.Body)
// 		}
// 	}()

// 	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
// 	<-forever
// }

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
