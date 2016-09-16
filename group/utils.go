package group

import (
	"github.com/soyking/douban-group-spider/proxy"
	"net/http"
	"net/url"
)

var (
	httpClient = &http.Client{
		// 使用一个不支持 302 跳转的 client，防止被删除帖子跳转到首页
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return ErrorTopicDelete
		},
	}
)

func SetProxy(b proxy.Balancer) {
	if b != nil {
		// 设置代理
		httpClient.Transport = &http.Transport{
			Proxy: func(r *http.Request) (*url.URL, error) {
				return b.Get(r), nil
			},
		}
	}
}
