package task

import (
	"github.com/soyking/douban-group-spider/filter"
	"github.com/soyking/douban-group-spider/flag"
	"github.com/soyking/douban-group-spider/group"
	"github.com/soyking/douban-group-spider/storage"
	"log"
	"strings"
	"sync"
	"time"
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

func NewTask(f *flag.Flag) (*Task, error) {
	filter, err := NewFilter(f)
	if err != nil {
		return nil, err
	}

	store, err := NewStorage(f)
	if err != nil {
		return nil, err
	}

	proxy, err := NewProxy(f)
	if err != nil {
		return nil, err
	}
	group.SetProxy(proxy)

	return &Task{
		groups:            strings.Split(f.GroupsName, ","),
		pages:             f.Pages,
		frequency:         f.Frequency,
		groupsConcurrency: f.GroupsConcurrency,
		topicsConcurrency: f.TopicsConcurrency,
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
