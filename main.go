package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coaraujo/go-vote-processor/stream"

	mongo "github.com/coaraujo/go-vote-processor/config/mongo"
	rabbit "github.com/coaraujo/go-vote-processor/config/rabbit"
	service "github.com/coaraujo/go-vote-processor/service"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting go-processor...")

	mongoCon := *mongo.NewConnection()
	rabbitCon := *rabbit.NewConnection()
	rabbitStream := *stream.NewRabbitStream(&rabbitCon, "go.vote")
	voteServ := service.NewVoteService(&mongoCon, &rabbitStream)

	rabbitStream.ReceiveVote()

	router := mux.NewRouter()
	router.HandleFunc("/vote", voteServ.SendVote).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", router))

	defer mongoCon.CloseConnection()
	defer rabbitCon.CloseConnection()
}
