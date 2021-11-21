package client

import (
	"github.com/soyking/douban-group-spider/refactor/fetcher"
)

const (
	host = "https://www.douban.com"
)

type Client struct {
	fetcher fetcher.Fetcher
}

func NewClient(fetcher fetcher.Fetcher) *Client {
	return &Client{
		fetcher: fetcher,
	}
}
