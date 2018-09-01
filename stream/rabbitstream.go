package stream

import (
	"fmt"
	"log"

	config "github.com/coaraujo/go-vote-processor/config/rabbit"
	"github.com/coaraujo/go-vote-processor/service"
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

	go func() {
		for vote := range votes {
			log.Printf("Received a message: %s", vote.Body)
			r.VoteService.SendVote2(vote.Body)
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
