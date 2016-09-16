package proxy

import (
	"net/http"
	"net/url"
	"sync"
)

type RoundRobinBalancer struct {
	sync.Mutex
	index int
	urls  []*url.URL
}

func (b *RoundRobinBalancer) Get(r *http.Request) *url.URL {
	b.Lock()
	defer b.Unlock()
	u := b.urls[b.index]
	b.index = (b.index + 1) % len(b.urls)
	return u
}

func NewRoundRobinProxy(urls []*url.URL) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		urls: urls,
	}
}
