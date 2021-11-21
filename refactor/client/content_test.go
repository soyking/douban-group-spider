package client

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/soyking/douban-group-spider/refactor/fetcher"
	"github.com/soyking/douban-group-spider/refactor/fetcher/impl/basic"
	"github.com/soyking/douban-group-spider/refactor/fetcher/impl/douban"
	"github.com/stretchr/testify/assert"
)

type TestFetcher struct {
	filepath string
}

func (f *TestFetcher) FetchURL(context.Context, string) (io.ReadCloser, error) {
	return os.Open(f.filepath)
}

func NewTestFetcher(filepath string) fetcher.Fetcher {
	return &TestFetcher{
		filepath: filepath,
	}
}
func TestClient_GetTopic_Local(t *testing.T) {
	// curl 'https://www.douban.com/group/topic/253060849/' > content_test.html
	client := NewClient(NewTestFetcher("content_test.html"))
	topicContent, err := client.GetTopic(context.TODO(), "")
	assert.Nil(t, err)

	b, _ := json.MarshalIndent(topicContent, "", "    ") // TODO: internal tools
	t.Logf("parsed topic content: %s", string(b))
}

func TestClient_GetTopic_Online(t *testing.T) {
	fetcher, err := basic.NewFetcher(
		basic.WithHTTPClient(douban.HTTPClient()),
		basic.WithRequestHandler(douban.RequestHandler()),
	)
	assert.Nil(t, err)

	client := NewClient(fetcher)
	topicContent, err := client.GetTopic(context.TODO(), "https://www.douban.com/group/topic/253060849/")
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(topicContent, "", "    ")
	t.Logf("parsed topic content: %s", string(b))
}
