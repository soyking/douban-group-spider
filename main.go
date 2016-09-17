package main

import (
	"github.com/soyking/douban-group-spider/flag"
	"github.com/soyking/douban-group-spider/group"
	"github.com/soyking/douban-group-spider/task"
	"log"
	"strings"
)

const (
	APP_NAME    = "DOUBAN GTOUP SPIDER"
	APP_VERSION = "0.0.1"
)

func main() {
	println(APP_NAME + "\t" + APP_VERSION)

	f := flag.ParseFlag()
	filter, err := newFilter(f)
	if err != nil {
		log.Fatal(err)
	}

	store, err := newStorage(f)
	if err != nil {
		log.Fatal(err)
	}

	proxy, err := newProxy(f)
	if err != nil {
		log.Fatal(err)
	}
	group.SetProxy(proxy)

	task, err := task.NewTask(
		strings.Split(f.GroupsName, ","),
		f.Pages,
		f.Frequency,
		f.GroupsConcurrency,
		f.TopicsConcurrency,
		filter,
		store,
	)
	if err != nil {
		log.Fatal(err)
	}

	task.Run()
}
