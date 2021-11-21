package proxy

import (
	"net/http"
	"net/url"
)

// proxy balancer for pick up proxy address
type Balancer interface {
	Get(*http.Request) (*url.URL, error)
}
