package main

import (
	"fmt"

	mongo "github.com/coaraujo/go-vote-processor/config/mongo"
	rabbit "github.com/coaraujo/go-vote-processor/config/rabbit"
	repository "github.com/coaraujo/go-vote-processor/repository"
	service "github.com/coaraujo/go-vote-processor/service"
	stream "github.com/coaraujo/go-vote-processor/stream"
)

func main() {
	fmt.Println("Starting go-processor...")

	mongoCon := mongo.NewConnection()
	rabbitCon := rabbit.NewConnection()

	groupRep := repository.NewGroupRepository()
	voteRep := repository.NewVoteRepository()

	voteServ := service.NewVoteService(voteRep, groupRep)
	groupServ := service.NewGroupService(voteRep, groupRep)
	groupServ.CreateFirstGroup()

	rabbitStream := *stream.NewRabbitStream(rabbitCon, voteServ, "go.vote")
	consumer := rabbitCon.CreateConsumerVote()
	rabbitStream.ListenVotes(consumer)

	defer mongoCon.CloseConnection()
	defer rabbitCon.CloseConnection()
}
