package es

import (
	"github.com/soyking/douban-group-spider/group"
	"gopkg.in/olivere/elastic.v3"
)

const (
	TOPIC_TYPE = "topic"
)

type ElasticSearchStorage struct {
	client *elastic.Client
	index  string
}

func NewElasticSearchStorage(addr, index string) (*ElasticSearchStorage, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://" + addr))
	if err != nil {
		return nil, err
	}

	exist, err := client.IndexExists(index).Do()
	if err != nil {
		return nil, err
	}

	if !exist {
		_, err := client.CreateIndex(index).Do()
		if err != nil {
			return nil, err
		}
		// mapping
		_, err = client.PutMapping().Index(index).Type(TOPIC_TYPE).BodyString(mappings).Do()
		if err != nil {
			return nil, err
		}
	}

	return &ElasticSearchStorage{
		client: client,
		index:  index,
	}, nil
}

func (e *ElasticSearchStorage) Save(topics []*group.Topic) error {
	bulkService := e.client.Bulk()
	for _, topic := range topics {
		id := topic.URL
		topic.URL = ""
		topicIndex := elastic.
			NewBulkIndexRequest().
			Index(e.index).
			Type(TOPIC_TYPE).
			Id(id).
			Doc(topic)
		bulkService.Add(topicIndex)
	}
	_, err := bulkService.Do()

	return err
}
