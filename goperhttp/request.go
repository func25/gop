package goperhttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func RequestJSON[T any](method string, url string, body []byte, opts ...RequestOption) (T, int, error) {
	var res T

	// config request
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return res, 0, err
	}

	myReq := httpRequest(*req)
	myReq.applyFrom(opts...)

	// send request
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return res, 0, err
	}

	// check status code
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return res, resp.StatusCode, err
	}

	// return body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, resp.StatusCode, err
	}

	err = json.Unmarshal(respBody, &res)
	return res, resp.StatusCode, err
}
