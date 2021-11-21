package proxy

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/soyking/douban-group-spider/refactor/fetcher/impl/basic"
)

func HTTPDynamicProxyClient(httpClient *http.Client, loader BalancerLoader) *http.Client {
	httpClient.Transport = &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			balancer, err := loader.Load(req)
			if err != nil {
				return nil, errors.Wrap(err, "load balancer")
			}

			if balancer != nil {
				if u, err := balancer.Get(req); err != nil {
					return nil, errors.Wrap(err, "get url from balancer")
				} else {
					return u, nil
				}
			}

			return nil, nil
		},
	}

	return httpClient
}

func NewHTTPClientFetcherOptionFunc(httpClient *http.Client, loader BalancerLoader) basic.FetcherOptionFunc {
	return basic.WithHTTPClient(
		HTTPDynamicProxyClient(httpClient, loader),
	)
}
