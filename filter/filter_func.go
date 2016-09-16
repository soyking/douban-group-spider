package filter

import (
	"github.com/soyking/douban-group-spider/group"
	"strings"
	"time"
)

type FilterFunc func(*group.Topic) bool

// 过滤中介用户等
func AuthorFilter(authors []string) FilterFunc {
	filterAuthors := []string{}
	filterAuthorURLs := []string{}
	for _, author := range authors {
		// 过滤用户可以是名字也可以是地址，豆瓣用户名不全站唯一
		if strings.HasPrefix(author, "http") {
			filterAuthorURLs = append(filterAuthorURLs, author)
		} else {
			filterAuthors = append(filterAuthors, author)
		}
	}

	return func(t *group.Topic) bool {
		for _, author := range filterAuthors {
			if t.Author == author {
				return false
			}
		}
		for _, authorURL := range filterAuthorURLs {
			if t.AuthorURL == authorURL {
				return false
			}
		}
		return true
	}
}

// 过滤包含特定字符串标题的帖子
func TitleFilter(titles []string) FilterFunc {
	return func(t *group.Topic) bool {
		for _, title := range titles {
			if strings.Contains(t.Title, title) {
				return false
			}
		}
		return true
	}
}

// 过滤包含特定字符串内容的帖子
func ContentFilter(contents []string) FilterFunc {
	return func(t *group.Topic) bool {
		if t.TopicContent == nil {
			return false
		}
		for _, content := range contents {
			if strings.Contains(t.TopicContent.Content, content) {
				return false
			}
		}
		return true
	}
}

// 回复数不能超过 maxReply 回复多可能是广告
func ReplyLimitFilter(maxReply int) FilterFunc {
	return func(t *group.Topic) bool {
		return t.Reply < maxReply
	}
}

// 必须有图
func PicFilter(withPic bool) FilterFunc {
	return func(t *group.Topic) bool {
		if t.TopicContent == nil {
			return false
		}
		return t.TopicContent.WithPic == withPic
	}
}

// 创建时间必须迟于 lastUpdateTime
func LastUpdateTimeFilter(lastUpdateTime time.Time) FilterFunc {
	return func(t *group.Topic) bool {
		if t.TopicContent == nil {
			return false
		}
		return t.TopicContent.UpdateTime.After(lastUpdateTime)
	}
}
