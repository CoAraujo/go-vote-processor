package service

import (
	"fmt"

	mongo "github.com/coaraujo/go-vote-processor/config/mongo"
	domain "github.com/coaraujo/go-vote-processor/domain"
)

type VoteService struct {
	MongoDB *mongo.MongoDB
}

func NewVoteService(m *mongo.MongoDB) *VoteService {
	voteService := VoteService{MongoDB: m}
	return &voteService
}

func (v *VoteService) SendVote(vote *domain.Vote) {
	fmt.Println("[VOTESERVICE] Send vote was invoked...")

	// _ = json.NewDecoder(r.Body).Decode(&vote)
}

func (v *VoteService) SendVote2(vote []byte) {
	fmt.Println("[VOTESERVICE] Send vote was invoked...")

	// _ = json.NewDecoder(r.Body).Decode(&vote)
}
