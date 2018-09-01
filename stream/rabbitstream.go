package stream

import (
	"encoding/json"
	"fmt"
	"log"

	config "github.com/coaraujo/go-vote-processor/config/rabbit"
	"github.com/coaraujo/go-vote-processor/domain"
	service "github.com/coaraujo/go-vote-processor/service"
	"github.com/streadway/amqp"
)

type RabbitStream struct {
	RabbitMQ    *config.RabbitMQ
	VoteService *service.VoteService
	Queue       amqp.Queue
	QueueGroup  string
}

func NewRabbitStream(conn *config.RabbitMQ, voteService *service.VoteService, queue string) *RabbitStream {
	q := "vote.groupBBB"
	if queue != "" {
		q = queue
	}
	rabbitStream := RabbitStream{RabbitMQ: conn, VoteService: voteService, QueueGroup: q}
	rabbitStream.Queue = conn.CreateQueue(q)
	return &rabbitStream
}

func (r *RabbitStream) ListenVotes(votes <-chan amqp.Delivery) {
	forever := make(chan bool)

	//Refazer aqui a goroutine.
	go func() {
		for vote := range votes {
			log.Printf("Received a message: %s", vote.Body)

			v := domain.Vote{}
			err := json.Unmarshal(vote.Body, &v)
			if err != nil {
				fmt.Println("There was an error:", err)
			}

			r.VoteService.SendVote(&v)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
