package task

import (
	"errors"
	"github.com/soyking/douban-group-spider/flag"
	"github.com/soyking/douban-group-spider/proxy"
	"net/url"
	"strings"
)

func NewProxy(f *flag.Flag) (proxy.Balancer, error) {
	if f.ProxyFile != flag.FLAG_PROXY_DEFAULT {
		addrsPorts, err := readLines(f.ProxyFile)
		if err != nil {
			return nil, err
		}

		proxyURLs := []*url.URL{}
		for _, line := range addrsPorts {
			addrPort := strings.Split(line, " ")
			if len(addrPort) != 2 {
				return nil, errors.New("proxy line format: addr port")
			}
			u, err := url.Parse("http://" + addrPort[0] + ":" + addrPort[1])
			if err != nil {
				return nil, err
			}
			proxyURLs = append(proxyURLs, u)
		}
		return proxy.NewRoundRobinProxy(proxyURLs), nil
	} else {
		return nil, nil
	}
}
