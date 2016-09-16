package group

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"os"
	"testing"
)

func getLocalContent() (*goquery.Document, error) {
	f, err := os.Open("./content_test.html")
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromReader(f)
}

func TestParseTopicContent(t *testing.T) {
	doc, err := getLocalContent()
	if err != nil {
		t.Error(err)
	} else {
		topicContent, err := ParseTopicContent(doc)
		if err != nil {
			t.Error(err)
		} else {
			b, _ := json.MarshalIndent(topicContent, "", "    ")
			t.Log(string(b))
		}
	}
}
