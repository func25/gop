package gopnet

import (
	"net/http"
	"time"
)

type httpRequest struct {
	*http.Request
	timeout       time.Duration
	acceptedCodes map[int]bool
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

func OptAcceptCode(statusCodes []int) RequestOption {
	return func(q *httpRequest) {
		q.acceptedCodes = make(map[int]bool, len(statusCodes))
		for _, v := range statusCodes {
			q.acceptedCodes[v] = true
		}
	}
}

func OptHeader(key string, value string) RequestOption {
	return func(q *httpRequest) {
		q.Header.Add(key, value)
	}
}
