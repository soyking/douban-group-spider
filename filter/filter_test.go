package filter

import (
	"github.com/soyking/douban-group-spider/group"
	"testing"
	"time"
)

func TestAuthorFilter(t *testing.T) {
	topic := &group.Topic{
		Author:    "soyking",
		AuthorURL: "https://www.douban.com/people/123/",
	}

	if AuthorFilter([]string{"soyking"})(topic) {
		t.Error("should not pass")
	} else if AuthorFilter([]string{"https://www.douban.com/people/123/"})(topic) {
		t.Error("should pass")
	} else if !AuthorFilter([]string{"soyking2"})(topic) {
		t.Error("should pass")
	}
}

func TestTitleFilter(t *testing.T) {
	topic := &group.Topic{
		Title: "中介是我",
	}

	if TitleFilter([]string{"嘿嘿", "中介"})(topic) {
		t.Error("should not pass")
	} else if !TitleFilter([]string{"nothing"})(topic) {
		t.Error("should pass")
	}
}

func TestContentFilter(t *testing.T) {
	topic := &group.Topic{
		TopicContent: &group.TopicContent{
			Content: "中介是我",
		},
	}

	if ContentFilter([]string{"嘿嘿", "中介"})(topic) {
		t.Error("should not pass")
	} else if !AuthorFilter([]string{"nothing"})(topic) {
		t.Error("should pass")
	}
}

func TestReplyLimitFilter(t *testing.T) {
	topic := &group.Topic{
		Reply: 100,
	}

	if ReplyLimitFilter(90)(topic) {
		t.Error("should not pass")
	} else if !ReplyLimitFilter(110)(topic) {
		t.Error("should pass")
	}
}

func TestPicFilter(t *testing.T) {
	topic := &group.Topic{
		TopicContent: &group.TopicContent{
			WithPic: true,
		},
	}

	if PicFilter(false)(topic) {
		t.Error("should not pass")
	} else if !PicFilter(true)(topic) {
		t.Error("should pass")
	}
}

func TestLastUpdateTimeFilter(t *testing.T) {
	now := time.Now()
	topic := &group.Topic{
		TopicContent: &group.TopicContent{
			UpdateTime: now,
		},
	}

	if LastUpdateTimeFilter(now.Add(time.Second))(topic) {
		t.Error("should not pass")
	} else if !LastUpdateTimeFilter(now.Add(-time.Second))(topic) {
		t.Error("should pass")
	}
}

func TestNewFilter(t *testing.T) {
	now := time.Now()
	filter := NewFilter(
		AuthorFilter([]string{"soyking"}),
		TitleFilter([]string{"中介"}),
		ContentFilter([]string{"中介"}),
		ReplyLimitFilter(100),
		PicFilter(true),
		LastUpdateTimeFilter(now),
	)
	topics := []*group.Topic{
		&group.Topic{
			Author: "soyking",
		},
		&group.Topic{
			Title: "中介是我",
		},
		&group.Topic{
			TopicContent: &group.TopicContent{
				Content: "中介是我",
			},
		},
		&group.Topic{
			Reply: 110,
		},
		&group.Topic{
			TopicContent: &group.TopicContent{
				WithPic: false,
			},
		},
		&group.Topic{
			TopicContent: &group.TopicContent{
				UpdateTime: now.Add(-time.Second),
			},
		},
		&group.Topic{
			Author: "soyking2",
			TopicContent: &group.TopicContent{
				WithPic:    true,
				UpdateTime: now.Add(time.Second),
			},
		},
	}
	topics = filter(topics)
	if len(topics) != 1 {
		t.Error("should only one pass")
	}
}
