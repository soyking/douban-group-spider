package proxy

import (
	"net/http"
	"net/url"
)

type Balancer interface {
	Get(*http.Request) *url.URL
}
