package group

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"os"
	"testing"
)

func getLocalGroup() (*goquery.Document, error) {
	f, err := os.Open("./group_test.html")
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return nil, err
	}

	doc.Url = &url.URL{}
	return doc, nil
}

func testTopics(t *testing.T, doc *goquery.Document) {
	topics, err := ParseTopics(doc)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("topics number: %d\n", len(topics))
		b, _ := json.MarshalIndent(topics, "", "    ")
		t.Log(string(b))
	}
}

func TestParseTopics(t *testing.T) {
	content, err := getLocalGroup()
	if err != nil {
		t.Error(err)
	} else {
		testTopics(t, content)
	}
}

func TestParseTopics2(t *testing.T) {
	content, err := GetGroup("beijingzufang")
	if err != nil {
		t.Error(err)
	} else {
		testTopics(t, content)
	}
}
