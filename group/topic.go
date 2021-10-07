package group

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const PAGE_TOPICS_NUMBER = 25

type Topic struct {
	URL           string        `json:"_id,omitempty" bson:"_id"` // 作为唯一键
	Title         string        `json:"title" bson:"title"`
	AuthorURL     string        `json:"author_url" bson:"author_url"`
	Author        string        `json:"author" bson:"author"`
	Reply         int           `json:"reply" bson:"reply"`
	LastReplyTime time.Time     `json:"last_reply_time" bson:"last_reply_time"`
	TopicContent  *TopicContent `json:"topic_content" bson:"topic_content"`
}

// name 为豆瓣小组名，有的是拼音如 beijingzufang，有的是数字
// pages 为每次获取页数，一页25条帖子
// cocurrency 为并发抓取帖子时的 goroutine 数量
func GetTopics(name string, pages int, concurrency ...int) ([]*Topic, error) {
	if pages < 1 {
		pages = 1
	}

	topics := []*Topic{}
	for i := 0; i < pages; i++ {
		log.Printf("\t\t[Info] try to fetch group %s page %d\n", name, i)
		doc, err := GetGroup(name, i*PAGE_TOPICS_NUMBER)
		if err != nil {
			return nil, err
		}

		pageTopics, err := ParseTopics(doc, concurrency...)
		if err != nil {
			return nil, err
		}
		topics = append(topics, pageTopics...)
	}

	return topics, nil
}

// 从文档树中获取 Topic
func ParseTopics(doc *goquery.Document, concurrency ...int) ([]*Topic, error) {
	var topics []*Topic
	var outErr []string
	nodes := doc.Find("#content > div > div.article > div:nth-child(2) > table > tbody > tr").Nodes

	con := 1
	if len(concurrency) > 0 {
		con = concurrency[0]
	}
	var wg sync.WaitGroup
	topicsChan := make(chan int, con)

	for i := range nodes {
		topicsChan <- 1
		wg.Add(1)
		go func(s *goquery.Selection) {
			topic, err := ParseTopic(s)
			if err != nil {
				outErr = append(outErr, "group: "+doc.Url.String()+" #"+strconv.Itoa(i)+"; "+err.Error())
			}

			if topic != nil {
				topics = append(topics, topic)
			}

			wg.Done()
			<-topicsChan
		}(&goquery.Selection{Nodes: []*html.Node{nodes[i]}})
	}

	wg.Wait()
	if len(outErr) != 0 {
		return topics, errors.New(strings.Join(outErr, " ## "))
	} else {
		return topics, nil
	}
}

// 解析成自定义的 Topic
func ParseTopic(s *goquery.Selection) (*Topic, error) {
	titleBlock := s.Find("td.title")
	if titleBlock.Length() == 0 {
		// 存在非小组话题
		return nil, nil
	}

	authorBlock := titleBlock.Next()
	replyBlock := authorBlock.Next()
	timeBlock := replyBlock.Next()

	titleBlock = titleBlock.Find("a")
	url, exist := titleBlock.Attr("href")
	if !exist {
		return nil, errors.New("without url")
	}
	topicContent, err := GetTopicContent(url)
	if err != nil {
		// http client 错误会包含其他
		if err == ErrorTopicDelete || strings.Contains(err.Error(), ErrorTopicDelete.Error()) {
			return nil, nil
		}
		return nil, errors.New("topic: " + url + "; " + err.Error())
	}

	title, exist := titleBlock.Attr("title")
	if !exist || title == "" {
		return nil, errors.New("without title")
	}

	authorBlock = authorBlock.Find("a")
	authorURL, exist := authorBlock.Attr("href")
	if !exist {
		return nil, errors.New("without author url")
	}
	author := authorBlock.Text()
	if author == "" {
		return nil, errors.New("without author")
	}

	replyStr := replyBlock.Text()
	reply := 0
	if replyStr != "" {
		var err error
		reply, err = strconv.Atoi(replyStr)
		if err != nil {
			return nil, err
		}
	}

	replyTimeStr := timeBlock.Text()
	if replyTimeStr == "" {
		return nil, errors.New("without last reply time")
	}
	// 时间格式是 08-31 23:42 加上当前年份方便解析成 time.Time
	lastReplyTime, err := time.Parse("2006-01-02 15:04", strconv.Itoa(time.Now().Year())+"-"+replyTimeStr)
	if err != nil {
		return nil, err
	}

	return &Topic{
		URL:           url,
		Title:         title,
		AuthorURL:     authorURL,
		Author:        author,
		Reply:         reply,
		LastReplyTime: lastReplyTime,
		TopicContent:  topicContent,
	}, nil
}
