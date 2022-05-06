package goperhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GoperHttpRes[T any] struct {
	Data       T
	StatusCode int
	Body       []byte
}

func RequestJSON[T any](method string, url string, body []byte, opts ...RequestOption) (GoperHttpRes[T], error) {
	var res GoperHttpRes[T]

	// config request
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return res, err
	}

	// apply custom
	myReq := httpRequest{
		Request:       req,
		timeout:       0,
		acceptedCodes: nil,
	}
	myReq.applyFrom(opts...)

	// send request
	client := &http.Client{
		Timeout: myReq.timeout,
	}
	resp, err := client.Do(myReq.Request)
	if err != nil {
		return GoperHttpRes[T]{}, err
	}

	// get body
	res.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	// get status code
	res.StatusCode = resp.StatusCode
	if len(myReq.acceptedCodes) != 0 {
		if !myReq.acceptedCodes[res.StatusCode] {
			return res, fmt.Errorf("wrong status code: %v", res.StatusCode)
		}
	} else if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return res, fmt.Errorf("wrong status code: %v", res.StatusCode)
	}

	// get data
	return res, json.Unmarshal(res.Body, &res.Data)
}
