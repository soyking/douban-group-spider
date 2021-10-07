package group

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrorForbidden = errors.New("request is forbidden, maybe too frequently")
)

func groupURL(name string, start ...int) string {
	s := "0"
	if len(start) > 0 {
		s = strconv.Itoa(start[0])
	}
	return host + "/group/" + name + "/discussion?start=" + s
}

// 获取豆瓣小组的内容 start 表示从第几条开始，默认 0
// 返回 25 条结果的网页内容，翻页时 start=25
func GetGroup(name string, start ...int) (*goquery.Document, error) {
	url := groupURL(name, start...)
	log.Printf("\t\t[Info] try to fetch url: %s\n", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request err: %w", err)
	}
	req = requestMiddleware(req)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusForbidden {
			return nil, ErrorForbidden
		}
		return nil, fmt.Errorf("response status %d err %w", resp.StatusCode, ErrorTopicDelete)
	}

	return goquery.NewDocumentFromReader(resp.Body)
}
