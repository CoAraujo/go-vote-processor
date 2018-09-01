package service

import (
	"fmt"
	"net/http"
	"time"

	domain "github.com/coaraujo/go-vote-processor/domain"
	repository "github.com/coaraujo/go-vote-processor/repository"
)

type GroupService struct {
	VoteRepository  *repository.VoteRepository
	GroupRepository *repository.GroupRepository
}

func NewGroupService(voteRepository *repository.VoteRepository, groupRepository *repository.GroupRepository) *GroupService {
	groupService := GroupService{VoteRepository: voteRepository, GroupRepository: groupRepository}
	return &groupService
}

func (v *GroupService) GetGroup(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("[VOTESERVICE] Send vote was invoked...")
	// var group domain.Group

}

func (v *GroupService) CreateFirstGroup() {
	fmt.Println("[GROUPSERVICE] Creating first group...")

	//TODO: Evita de criar quando já tiver sido criado.
	group := domain.Group{ID: "1", CreatedDate: time.Now().UTC(), EndTime: time.Now().UTC().AddDate(0, 0, 1), Options: []int{1, 2}}

	v.GroupRepository.InsertGroup(&group)
}
