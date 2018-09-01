package config

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const (
	rabbitURL = "amqp://guest:guest@localhost:5672"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

func (r *RabbitMQ) GetChannel() *amqp.Channel {
	return r.Channel
}

func NewConnection() *RabbitMQ {
	rabbitmq := RabbitMQ{}

	rabbitmq.getConnection()
	rabbitmq.setChannel()

	return &rabbitmq
}

func (r *RabbitMQ) CloseConnection() {
	r.Connection.Close()
	r.Channel.Close()
}

func (r *RabbitMQ) CreateQueue(queueGroup string) amqp.Queue {
	q, err := r.Channel.QueueDeclare(
		queueGroup, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "[RABBITMQ] Failed to declare a queue.")
	r.Queue = q
	return r.Queue
}

func (r *RabbitMQ) CreateConsumerVote() <-chan amqp.Delivery {
	fmt.Println("[RABBITMQ] Receiving vote...")

	votes, err := r.Channel.Consume(
		r.Queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")
	return votes
}

func (r *RabbitMQ) getConnection() {
	fmt.Println("[RABBITMQ] Connecting...")

	conn, err := amqp.Dial(rabbitURL)
	failOnError(err, "[RABBITMQ] Failed to connect to RabbitMQ")
	r.Connection = conn

	fmt.Println("[RABBITMQ] Connected sucessfully.")
}

func (r *RabbitMQ) setChannel() {
	fmt.Println("[RABBITMQ] Setting channel...")

	ch, err := r.Connection.Channel()
	failOnError(err, "[RABBITMQ] Failed to open a channel")
	r.Channel = ch

	fmt.Println("[RABBITMQ] Setup channel successfully.")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
