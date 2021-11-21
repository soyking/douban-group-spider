package fetcher

import (
	"context"
	"io"
)

type Fetcher interface {
	FetchURL(ctx context.Context, url string) (io.ReadCloser, error)
}
