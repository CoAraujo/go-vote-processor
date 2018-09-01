package service

import (
	"net/http"

	mongo "github.com/coaraujo/go-vote-processor/config/mongo"
)

type GroupService struct {
	MongoDB *mongo.MongoDB
}

func NewGroupService(m *mongo.MongoDB) *GroupService {
	groupService := GroupService{MongoDB: m}
	return &groupService
}

func (v *GroupService) GetGroup(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("[VOTESERVICE] Send vote was invoked...")
	// var group domain.Group

}
