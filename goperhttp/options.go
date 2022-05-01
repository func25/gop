package goperhttp

import "net/http"

type httpRequest http.Request

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
