package repository

import (
	"fmt"
	"log"

	config "github.com/coaraujo/go-vote-processor/config/mongo"
	domain "github.com/coaraujo/go-vote-processor/domain"
	"gopkg.in/mgo.v2/bson"
)

const (
	paredaoDatabase       = "vote"
	paredaoCollectionName = "paredao"
)

type ParedaoRepository struct {
	MongoDB *config.MongoDB
}

func newParedaoRepository() *ParedaoRepository {
	paredaoRepository := ParedaoRepository{MongoDB: config.NewConnection()}
	return &paredaoRepository
}

func (p *ParedaoRepository) InsertParedao(paredao domain.Paredao) {
	fmt.Println("[MONGODB] Inserting value: ", paredao)

	err := p.MongoDB.GetCollection(paredaoDatabase, paredaoCollectionName).Insert(paredao)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Value inserted: ", paredao)
}

func (p *ParedaoRepository) GetParedaoById(paredaoId string) {
	fmt.Println("[MONGODB] Getting paredao for paredaoId:", paredaoId)

	result := domain.Paredao{}
	err := p.MongoDB.GetCollection(paredaoDatabase, paredaoCollectionName).Find(bson.M{"id": paredaoId}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[MONGODB] Get paredao:", result)
}
