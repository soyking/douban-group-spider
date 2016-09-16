package group

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

var (
	ErrorForbidden = errors.New("request is forbidden, maybe too frequently")
)

func getGroupURL(name string, start ...int) string {
	s := "0"
	if len(start) > 0 {
		s = strconv.Itoa(start[0])
	}
	return "https://www.douban.com/group/" + name + "/discussion?start=" + s
}

// 获取豆瓣小组的内容 start 表示从第几条开始 默认 0
// 返回 25 条结果的网页内容
func GetGroup(name string, start ...int) (*goquery.Document, error) {
	resp, err := httpClient.Get(getGroupURL(name, start...))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusForbidden {
			return nil, ErrorForbidden
		}
		return nil, ErrorTopicDelete
	}

	return goquery.NewDocumentFromResponse(resp)
}
