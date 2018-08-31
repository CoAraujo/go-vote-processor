package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

const (
	host = "localhost"
)

type MongoDB struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func (m *MongoDB) GetCollection(database string, collectionName string) *mgo.Collection {
	c := m.session.DB(database).C(collectionName)
	return c
}

func NewConnection() *MongoDB {
	fmt.Println("[MONGODB] Connecting...")

	mongodb := MongoDB{}
	s, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	mongodb.session = s

	fmt.Println("[MONGODB] Connection started sucessfully.")
	return &mongodb
}

func (m *MongoDB) CloseConnection() {
	fmt.Println("[MONGODB] Closing connection...")
	m.session.Close()
	fmt.Println("[MONGODB] Connection closed.")
}
