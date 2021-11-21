package proxy

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/pkg/errors"
)

type RoundRobinBalancer struct {
	lock sync.Mutex

	index int
	urls  []*url.URL
}

func (b *RoundRobinBalancer) Get(r *http.Request) (*url.URL, error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	u := b.urls[b.index]
	b.index = (b.index + 1) % len(b.urls)
	return u, nil
}

func NewRoundRobinProxy(urls []*url.URL) (Balancer, error) {
	if len(urls) == 0 {
		return nil, errors.New("without urls")
	}
	return &RoundRobinBalancer{
		urls:  urls,
		index: 0,
	}, nil
}
