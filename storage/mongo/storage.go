package mongo

import (
	"github.com/soyking/douban-group-spider/group"
	"gopkg.in/mgo.v2/bson"
)

const (
	TOPIC_COLLECTION = "topic"
)

type MongoDBStorage struct {
	mongoDBHandler *MongoDBHandler
}

func NewMongoDBStorage(addr, username, password, database string) (*MongoDBStorage, error) {
	m, err := NewMongoDBHandler(addr, username, password, database, TOPIC_COLLECTION)
	if err != nil {
		return nil, err
	}
	return &MongoDBStorage{m}, nil
}

func (m *MongoDBStorage) Save(topics []*group.Topic) error {
	pairs := []interface{}{}
	for i := range topics {
		pairs = append(pairs, bson.M{"_id": topics[i].URL}, topics[i])
	}
	_, err := m.mongoDBHandler.BulkUpsert(pairs...)
	return err
}
