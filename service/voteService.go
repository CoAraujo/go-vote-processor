package service

import (
	"fmt"
	"time"

	domain "github.com/coaraujo/go-vote-processor/domain"
	repository "github.com/coaraujo/go-vote-processor/repository"
)

type VoteService struct {
	VoteRepository  *repository.VoteRepository
	GroupRepository *repository.GroupRepository
}

func NewVoteService(voteRepository *repository.VoteRepository, groupRepository *repository.GroupRepository) *VoteService {
	voteService := VoteService{VoteRepository: voteRepository, GroupRepository: groupRepository}
	return &voteService
}

func (v *VoteService) PersistVote(vote *domain.Vote) {
	fmt.Println("[VOTESERVICE] Send vote was invoked...")

	// Validar se existe esse groupid
	group := v.GroupRepository.GetGroupById(vote.GroupID)
	fmt.Println("[VOTESERVICE] GROUP WAS: ", group)

	// Validar se este group ainda está aberto, ou seja, se a data de agora é menor que a data de finalizacao
	if group.EndTime.Before(time.Now()) {
		fmt.Println("[VOTESERVICE]INVALID VOTE! GROUP IS CLOSED.")
		return
	}

	//Validar se neste group, podem estas options
	for option, _ := range group.Options {

		//Se for true para tudo, persistir no banco.
		if vote.Option == option {
			v.VoteRepository.InsertVote(vote)
		}
	}

	fmt.Println("[VOTESERVICE] VOTE WAS: ", vote)

}
