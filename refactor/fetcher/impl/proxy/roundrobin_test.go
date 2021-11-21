package proxy

import (
	"net/url"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundRobinBalancer(t *testing.T) {
	for _, urlLen := range []int{1, 2, 4, 7, 16, 101, 1024, 10000} {
		for _, concurrency := range []int{urlLen, urlLen - 1, urlLen + 1, urlLen * 2} {
			var urls []*url.URL
			for i := 0; i < urlLen; i++ {
				urls = append(urls, &url.URL{})
			}
			balancer, err := NewRoundRobinProxy(urls)
			assert.Nil(t, err)

			var wg sync.WaitGroup
			wg.Add(concurrency)
			for i := 0; i < concurrency; i++ {
				go func() {
					defer wg.Done()
					_, err := balancer.Get(nil)
					assert.Nil(t, err)
				}()
			}

			wg.Wait()
			nextProxy, err := balancer.Get(nil)
			assert.Nil(t, err)
			assert.Same(t, urls[concurrency%urlLen], nextProxy)
		}
	}
}
