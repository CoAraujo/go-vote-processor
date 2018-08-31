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

func newVoteRepository() *VoteRepository {
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

func (v *VoteRepository) GetVote(paredaoId string) {
	fmt.Println("[MONGODB] Getting vote for paredaoId:", paredaoId)

	result := domain.Vote{}
	err := v.MongoDB.GetCollection(voteDatabase, voteCollectionName).Find(bson.M{"name": paredaoId}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Get vote:", result)
}
