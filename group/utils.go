package group

import (
	"net/http"
	"net/url"
	"time"

	"github.com/soyking/douban-group-spider/proxy"
)

var (
	httpClient = &http.Client{
		// 使用一个不支持 302 跳转的 client，防止被删除帖子跳转到首页
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return ErrorTopicDelete
		},
		Timeout: 5 * time.Second,
	}
)

func requestMiddleware(req *http.Request) *http.Request {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	return req
}

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
