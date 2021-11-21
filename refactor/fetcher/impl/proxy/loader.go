package proxy

import (
	"net/http"
	"sync"
)

type BalancerLoader interface {
	Load(*http.Request) (Balancer, error)
}

type BasicBalancerLoader struct {
	rwLock   sync.RWMutex
	balancer Balancer
}

func (l *BasicBalancerLoader) Load(*http.Request) (Balancer, error) {
	l.rwLock.RLock()
	defer l.rwLock.RUnlock()
	return l.balancer, nil
}

func (l *BasicBalancerLoader) UpdateBalancer(balancer Balancer) {
	l.rwLock.Lock()
	defer l.rwLock.Unlock()
	l.balancer = balancer
}

func NewBasicBalancerLoader(balancer Balancer) BalancerLoader {
	return &BasicBalancerLoader{
		balancer: balancer,
	}
}
