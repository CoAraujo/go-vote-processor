package service

import (
	"fmt"

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

func (v *VoteService) SendVote(vote *domain.Vote) {
	fmt.Println("[VOTESERVICE] Send vote was invoked...")

	//Validar se existe esse groupid
	group := v.GroupRepository.GetGroupById(vote.GroupID)
	fmt.Println("[VOTESERVICE] GROUP WAS: ", group)

	//Validar se este group ainda está aberto, ou seja, se a data de agora é menor que a data de finalizacao

	//Validar se neste group, podem estas options

	//Se for true para tudo, persistir no banco.

	fmt.Println("[VOTESERVICE] VOTE WAS: ", &vote)

}
