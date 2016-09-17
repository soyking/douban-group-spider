package main

import (
	"github.com/soyking/douban-group-spider/flag"
	"github.com/soyking/douban-group-spider/storage"
	"github.com/soyking/douban-group-spider/storage/es"
	"github.com/soyking/douban-group-spider/storage/mongo"
)

func newStorage(f *flag.Flag) (storage.StorageSave, error) {
	var store storage.StorageSave
	var err error
	if f.MongoDBOn {
		println("STORAGE MONGODB ( " + f.MongoDBAddr + " )\n")
		store, err = mongo.NewMongoDBStorage(f.MongoDBAddr, f.MongoDBUsername, f.MongoDBPassword, f.MongoDBDatabase)
	} else {
		println("STORAGE ELASTICSEARCH ( " + f.EsAddr + " )\n")
		store, err = es.NewElasticSearchStorage(f.EsAddr, f.EsIndex)
	}
	return store, err
}
