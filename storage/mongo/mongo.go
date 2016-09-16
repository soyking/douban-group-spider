package mongo

import (
	"gopkg.in/mgo.v2"
	"strings"
)

type MongoDBHandler struct {
	database   string
	collection string
	session    *mgo.Session
}

func NewMongoDBHandler(addr, username, password, database, collection string) (*MongoDBHandler, error) {
	dialInfo := &mgo.DialInfo{
		Addrs:    strings.Split(addr, ","),
		Username: username,
		Password: password,
		Database: database,
	}
	sess, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}

	err = sess.Ping()
	if err != nil {
		return nil, err
	}

	return &MongoDBHandler{
		database:   database,
		collection: collection,
		session:    sess,
	}, nil
}

func (m *MongoDBHandler) BulkUpsert(pairs ...interface{}) (*mgo.BulkResult, error) {
	sess := m.session.Copy()
	defer sess.Close()
	bulk := sess.DB(m.database).C(m.collection).Bulk()
	bulk.Upsert(pairs...)
	return bulk.Run()
}
