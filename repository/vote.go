package repository

import (
	"fmt"
	"log"

	mongo "github.com/coaraujo/go-vote-processor/config/mongo"
	domain "github.com/coaraujo/go-vote-processor/domain"
	"gopkg.in/mgo.v2/bson"
)

const (
	voteDatabase       = "vote"
	voteCollectionName = "vote"
)

type VoteRepository struct {
	MongoDB *mongo.MongoDB
}

func NewVoteRepository() *VoteRepository {
	voteRepository := VoteRepository{MongoDB: mongo.NewConnection()}
	return &voteRepository
}

func (v *VoteRepository) InsertVote(vote domain.Vote) {
	fmt.Println("[MONGODB] Inserting value: ", vote)

	err := v.MongoDB.GetCollection(voteDatabase, voteCollectionName).Insert(vote)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Value inserted: ", vote)
}

func (v *VoteRepository) GetVoteByGroupId(groupId string) {
	fmt.Println("[MONGODB] Getting vote for groupId:", groupId)

	result := domain.Vote{}
	err := v.MongoDB.GetCollection(voteDatabase, voteCollectionName).Find(bson.M{"groupId": groupId}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Get vote:", result)
}
