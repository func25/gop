package gopnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

type NetRes[T any] struct {
	Data T
	Code int
	Body []byte
}

func Call[T any](method string, url string, body interface{}, opts ...RequestOption) (res NetRes[T], err error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return
	}

	// config request
	req, err := http.NewRequest(method, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return
	}

	// apply custom
	myReq := httpRequest{
		Request:       req,
		timeout:       0,
		acceptedCodes: nil,
	}

	if body != nil {
		opts = append([]RequestOption{OptHeader("Content-Type", "application/json;charset=utf-8")}, opts...)
	}

	myReq.applyFrom(opts...)

	// send request
	client := &http.Client{
		Timeout: myReq.timeout,
	}
	resp, err := client.Do(myReq.Request)
	if err != nil {
		return
	}

	// get body
	res.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// get status code
	res.Code = resp.StatusCode
	if len(myReq.acceptedCodes) != 0 {
		if !myReq.acceptedCodes[res.Code] {
			return res, fmt.Errorf("wrong status code: %v", res.Code)
		}
	} else if res.Code < http.StatusOK || res.Code >= http.StatusBadRequest {
		return res, fmt.Errorf("wrong status code: %v", res.Code)
	}

	// get data
	return res, json.Unmarshal(res.Body, &res.Data)
}

func Get[T any](url string, queryStr interface{}, opts ...RequestOption) (res NetRes[T], err error) {
	v, err := query.Values(queryStr)
	if err != nil {
		return res, err
	}

	if len(v) > 0 {
		url = fmt.Sprintf("%s?%s", url, v.Encode())
	}
	return Call[T]("GET", url, nil, opts...)
}

func Post[T any](url string, body interface{}, opts ...RequestOption) (NetRes[T], error) {
	return Call[T]("POST", url, body, opts...)
}

func Put[T any](url string, body interface{}, opts ...RequestOption) (NetRes[T], error) {
	return Call[T]("PATCH", url, body, opts...)
}

func Patch[T any](url string, body interface{}, opts ...RequestOption) (NetRes[T], error) {
	return Call[T]("PATCH", url, body, opts...)
}
