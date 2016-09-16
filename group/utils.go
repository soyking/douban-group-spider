package group

import "net/http"

var (
	httpClient = &http.Client{
		// 使用一个不支持 302 跳转的 client，防止被删除帖子跳转到首页
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return ErrorTopicDelete
		},
	}
)
