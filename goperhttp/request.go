package goperhttp

import (
	"bytes"
	"encoding/json"
	"errors"
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
	if (len(myReq.acceptedCodes) != 0 && !myReq.acceptedCodes[res.StatusCode]) ||
		res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return res, errors.New("[goper] wrong status code: ")
	}

	// get data
	return res, json.Unmarshal(res.Body, &res.Data)
}
