package basic

import (
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/soyking/douban-group-spider/refactor/fetcher"
)

type RequestHandlerFunc func(ctx context.Context, req *http.Request) error
type ResponseHandlerFunc func(ctx context.Context, resp *http.Response) (io.ReadCloser, error)

type Fetcher struct {
	httpClient      *http.Client
	requestHandler  RequestHandlerFunc
	responseHandler ResponseHandlerFunc
}

func (c *Fetcher) FetchURL(ctx context.Context, url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new request")
	}

	if c.requestHandler != nil {
		if err := c.requestHandler(ctx, req); err != nil {
			return nil, errors.Wrap(err, "handle request")
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do request")
	}

	if stream, err := c.responseHandler(ctx, resp); err != nil {
		return nil, errors.Wrap(err, "handle response")
	} else {
		return stream, nil
	}
}

func NewFetcher(optionFuncs ...FetcherOptionFunc) (fetcher.Fetcher, error) {
	options := defaultFetcherOptions()
	for _, optionFunc := range optionFuncs {
		optionFunc(options)
	}

	fetcher := &Fetcher{
		httpClient:      options.httpClient,
		requestHandler:  options.requestHandler,
		responseHandler: options.responseHandler,
	}

	if fetcher.httpClient == nil {
		return nil, errors.New("without http client")
	}

	if fetcher.responseHandler == nil {
		return nil, errors.New("without response handler")
	}

	return fetcher, nil
}

/* fetcher options */

type fetcherOptions struct {
	httpClient      *http.Client
	requestHandler  RequestHandlerFunc
	responseHandler ResponseHandlerFunc
}

type FetcherOptionFunc func(options *fetcherOptions)

func WithHTTPClient(httpClient *http.Client) FetcherOptionFunc {
	return func(options *fetcherOptions) {
		options.httpClient = httpClient
	}
}

func WithRequestHandler(requestHandler RequestHandlerFunc) FetcherOptionFunc {
	return func(options *fetcherOptions) {
		options.requestHandler = requestHandler
	}
}

func WithResponseHandler(responseHandler ResponseHandlerFunc) FetcherOptionFunc {
	return func(options *fetcherOptions) {
		options.responseHandler = responseHandler
	}
}

func DefaultHTTPClient() *http.Client {
	return &http.Client{}
}

func DefaultResponseHandler() ResponseHandlerFunc {
	return func(c context.Context, resp *http.Response) (io.ReadCloser, error) {
		return resp.Body, nil
	}
}

func defaultFetcherOptions() *fetcherOptions {
	return &fetcherOptions{
		httpClient:      DefaultHTTPClient(),
		responseHandler: DefaultResponseHandler(),
	}
}
