package repository

import (
	"fmt"
	"log"

	config "github.com/coaraujo/go-vote-processor/config/mongo"
	domain "github.com/coaraujo/go-vote-processor/domain"
	"gopkg.in/mgo.v2/bson"
)

const (
	groupDatabase       = "vote"
	groupCollectionName = "group"
)

type GroupRepository struct {
	MongoDB *config.MongoDB
}

func NewGroupRepository() *GroupRepository {
	groupRepository := GroupRepository{MongoDB: config.NewConnection()}
	return &groupRepository
}

func (p *GroupRepository) InsertGroup(group domain.Group) {
	fmt.Println("[MONGODB] Inserting value: ", group)

	err := p.MongoDB.GetCollection(groupDatabase, groupCollectionName).Insert(group)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Value inserted: ", group)
}

func (p *GroupRepository) GetGroupById(id string) *domain.Group {
	fmt.Println("[MONGODB] Getting group by id:", id)

	result := domain.Group{}
	err := p.MongoDB.GetCollection(groupDatabase, groupCollectionName).Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Group:", result)
	return &result
}
