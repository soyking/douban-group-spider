package main

import (
	"github.com/soyking/douban-group-spider/flag"
	"github.com/soyking/douban-group-spider/task"
	"log"
)

const (
	APP_NAME    = "DOUBAN RENT TOOLS - SPIDER"
	APP_VERSION = "0.0.1"
)

func main() {
	println(APP_NAME + "\t" + APP_VERSION)

	f := flag.ParseFlag()
	task, err := task.NewTask(f)
	if err != nil {
		log.Fatal(err)
	}

	task.Run()
}
