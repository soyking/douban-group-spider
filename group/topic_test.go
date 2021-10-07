package group

import (
	"encoding/json"
	"net/url"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func getLocalGroup() (*goquery.Document, error) {
	// build group_test.html:
	//     curl 'https://www.douban.com/group/beijingzufang/discussion?start=0' > group/group_test.html
	f, err := os.Open("./group_test.html")
	if err != nil {
		return nil, err
	}
	defer f.Close()

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
		t.Fatal(err)
	}

	if len(topics) == 0 {
		t.Fatal("could not fetch topics")
	}

	t.Logf("topics number: %d\n", len(topics))
	b, _ := json.MarshalIndent(topics, "", "    ")
	t.Logf("fetched topics: %s", string(b))
}

func TestParseTopicsLocal(t *testing.T) {
	content, err := getLocalGroup()
	if err != nil {
		t.Fatal(err)
	}

	testTopics(t, content)
}

func TestParseTopicsRemote(t *testing.T) {
	content, err := GetGroup("beijingzufang")
	if err != nil {
		t.Error(err)
	}

	testTopics(t, content)
}
