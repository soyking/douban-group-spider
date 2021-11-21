package proxy

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func HTTPDynamicProxyClient(httpClient *http.Client, loader BalancerLoader) (*http.Client, error) {
	if transport, ok := httpClient.Transport.(*http.Transport); ok {
		transport.Proxy = func(req *http.Request) (*url.URL, error) {
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
		}
		return httpClient, nil
	} else {
		return nil, errors.New("unknown http.Transport")
	}
}
