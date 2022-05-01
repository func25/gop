package goperhttp

import (
	"net/http"
	"time"
)

type httpRequest struct {
	*http.Request
	timeout time.Duration
}

type RequestOption func(*httpRequest)

func (c *httpRequest) applyFrom(opts ...RequestOption) *httpRequest {
	for i := range opts {
		opts[i](c)
	}

	return c
}

func OptAuthorization(token string) RequestOption {
	return func(q *httpRequest) {
		q.Header.Add("Authorization", token)
	}
}

func OptTimeout(timeout time.Duration) RequestOption {
	return func(q *httpRequest) {
		q.timeout = timeout
	}
}
