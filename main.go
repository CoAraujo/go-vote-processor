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
	rabbitStream := *stream.NewRabbitStream(&rabbitCon)
	voteServ := service.NewVoteService(&mongoCon, &rabbitStream)

	router := mux.NewRouter()
	router.HandleFunc("/vote", voteServ.SendVote).Methods("POST")
	// router.HandleFunc("/vote/{id}", voteServ.GetVote).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

	defer mongoCon.CloseConnection()
	defer rabbitCon.CloseConnection()
}
