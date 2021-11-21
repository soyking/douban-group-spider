package fetcher

import "context"

type Fetcher interface {
	FetchURL(ctx context.Context, url string) (string, error)
}
