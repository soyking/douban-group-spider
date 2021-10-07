package group

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func getLocalContent() (*goquery.Document, error) {
	// build group_test.html:
	//     curl 'https://www.douban.com/group/topic/249109397/' > group/content_test.html
	f, err := os.Open("./content_test.html")
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromReader(f)
}

func TestParseTopicContent(t *testing.T) {
	doc, err := getLocalContent()
	if err != nil {
		t.Fatal(err)
	}

	topicContent, err := ParseTopicContent(doc)
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(topicContent, "", "    ")
	t.Logf("parsed topic content: %s", string(b))
}
