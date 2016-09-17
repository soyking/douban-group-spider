package task

import (
	"errors"
	"github.com/soyking/douban-group-spider/filter"
	"github.com/soyking/douban-group-spider/group"
	"github.com/soyking/douban-group-spider/storage"
	"log"
	"sync"
	"time"
)

var (
	ErrorNoGroups = errors.New("no groups")
	ErrorNoFilter = errors.New("no filter")
	ErrorNoStore  = errors.New("no store")
)

type Task struct {
	groups            []string
	pages             int
	frequency         int64
	groupsConcurrency int
	topicsConcurrency int
	filter            filter.Filter
	store             storage.StorageSave
}

func NewTask(
	groups []string,
	pages int,
	frequency int64,
	groupsConcurrency int,
	topicsConcurrency int,
	filter filter.Filter,
	store storage.StorageSave,
) (*Task, error) {
	if len(groups) == 0 {
		return nil, ErrorNoGroups
	}

	if filter == nil {
		return nil, ErrorNoFilter
	}

	if store == nil {
		return nil, ErrorNoStore
	}

	return &Task{
		groups:            groups,
		pages:             pages,
		frequency:         frequency,
		groupsConcurrency: groupsConcurrency,
		topicsConcurrency: topicsConcurrency,
		filter:            filter,
		store:             store,
	}, nil
}

func (t *Task) Run() {
	log.Printf("crawling groups: %s\n", t.groups)
	log.Println("...start task...")

	tick := time.Tick(time.Duration(t.frequency) * time.Second)
	count := 1

	for _ = range tick {
		log.Printf("\ttask %d\n", count)

		var wg sync.WaitGroup
		taskChan := make(chan int, t.groupsConcurrency)

		for _, g := range t.groups {
			taskChan <- 1
			wg.Add(1)

			go func(groupName string) {
				topics, err := group.GetTopics(groupName, t.pages, t.topicsConcurrency)
				if err != nil {
					log.Printf("\t\t[Fail] fetch group: %s err: %s\n", groupName, err.Error())
				}

				if len(topics) != 0 {
					topics = t.filter(topics)
					err = t.store.Save(topics)
					if err != nil {
						log.Printf("\t\t[Fail] save group: %s err: %s\n", groupName, err.Error())
					} else {
						log.Printf("\t\t[SUCCESS] group: %s topics %d\n", groupName, len(topics))
					}
				}

				wg.Done()
				<-taskChan
			}(g)
		}

		wg.Wait()
		count += 1
	}
}
