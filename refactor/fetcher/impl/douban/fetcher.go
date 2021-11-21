package douban

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/soyking/douban-group-spider/refactor/fetcher/impl/basic"
)

var (
	ErrorRedirect = errors.New("redirect happens")
)

func HTTPClient() *http.Client {
	return &http.Client{
		// 使用一个不支持 302 跳转的 client，防止被删除帖子跳转到首页
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return ErrorRedirect
		},
		Timeout: time.Minute,
	}
}

func RequestHandler() basic.RequestHandlerFunc {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		return nil
	}
}
