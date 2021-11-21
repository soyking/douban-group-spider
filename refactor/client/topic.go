package client

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

var (
	// spec errors
	ErrorTopicDelete = errors.New("topic has been deleted")

	emptyWords = []string{
		" ",
		"\r",
		"\n",
		"\t",
	}
	emptyReplacer *strings.Replacer
)

func init() {
	var oldnew []string
	for _, w := range emptyWords {
		// 替换为空
		oldnew = append(oldnew, w, "")
	}
	emptyReplacer = strings.NewReplacer(oldnew...)
}

type Topic struct {
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
	Content    string    `json:"content,omitempty" bson:"content"`
	WithPic    bool      `json:"with_pic" bson:"with_pic"`
	PicURLs    []string  `json:"pic_urls,omitempty" bson:"pic_urls"`
	Like       int       `json:"like" bson:"like"`
}

func (c *Client) GetTopic(ctx context.Context, url string) (*Topic, error) {
	stream, err := c.fetcher.FetchURL(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "fetch url")
	}
	defer stream.Close()

	doc, err := goquery.NewDocumentFromReader(stream)
	if err != nil {
		return nil, errors.Wrap(err, "go query new document")
	}

	topicContent, err := ParseTopicContent(doc)
	if err != nil {
		return nil, errors.Wrapf(err, "parse topic content: %s", doc.Text())
	}

	return topicContent, nil
}

func ParseTopicContent(doc *goquery.Document) (*Topic, error) {
	updateTimeStr := doc.Find("#topic-content > div.topic-doc > h3 > span.create-time.color-green").Text()
	if updateTimeStr == "" {
		// 存在蓝色状态的帖子，感觉是一种预删除的状态，页面结构不一样，这里作帖子被删除处理
		return nil, ErrorTopicDelete
	}

	updateTime, err := time.Parse("2006-01-02 15:04:05", updateTimeStr)
	if err != nil {
		return nil, err
	}

	topicBlock := doc.Find("#link-report > div")
	if topicBlock.Length() == 0 {
		return nil, errors.New("without content")
	}

	content := []string{}
	topicBlock.Find("p").Each(func(i int, s *goquery.Selection) {
		content = append(content, emptyReplacer.Replace(s.Text()))
	})
	picBlock := topicBlock.Find("#link-report > div > div > div")
	withPic := false
	picURLs := []string{}
	if picBlock.Length() > 0 {
		withPic = true
		picBlock.Each(func(i int, s *goquery.Selection) {
			picURL, exist := s.Find("img").Attr("src")
			if exist && picURL != "" {
				picURLs = append(picURLs, picURL)
			}
		})
	}

	likeStr := doc.Find("#sep > div.action-react > a > span.react-num").Text()
	like := 0
	if likeStr != "" {
		like, err = strconv.Atoi(strings.TrimSpace(likeStr))
		if err != nil {
			return nil, err
		}
	}

	return &Topic{
		UpdateTime: updateTime,
		Content:    strings.Join(content, ""),
		WithPic:    withPic,
		PicURLs:    picURLs,
		Like:       like,
	}, nil
}
