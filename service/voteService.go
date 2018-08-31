package voteService

import (
	"encoding/json"
	"fmt"
	"net/http"

	mongo "github.com/coaraujo/go-vote-processor/config/mongo"
	domain "github.com/coaraujo/go-vote-processor/domain"
	stream "github.com/coaraujo/go-vote-processor/stream"
)

type VoteService struct {
	MongoDB      *mongo.MongoDB
	RabbitStream *stream.RabbitStream
}

func NewVoteService(m *mongo.MongoDB, s *stream.RabbitStream) *VoteService {
	voteService := VoteService{MongoDB: m, RabbitStream: s}
	return &voteService
}

func (v *VoteService) SendVote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[VOTESERVICE] Send vote was invoked...")

	var vote domain.Vote
	_ = json.NewDecoder(r.Body).Decode(&vote)
}
