package basic

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetcher(t *testing.T) {
	tsResponseBody := "tsResponseBody"
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		_, _ = rw.Write([]byte(tsResponseBody))
	}))
	defer ts.Close()

	fetcher, err := NewFetcher()
	assert.Nil(t, err)

	data, err := fetcher.FetchURL(context.TODO(), ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, data, tsResponseBody)
}
