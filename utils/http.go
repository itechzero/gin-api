package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type RequestErr struct {
	Err   error       `json:"err"`
	Uri   string      `json:"uri"`
	Param interface{} `json:"param"`
	Data  interface{} `json:"data"`
}

//GET METHOD
func HttpGet(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	// new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("new request is fail ")
	}

	// 拼接Url请求参数
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	// 拼接header参数
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	client := &http.Client{}

	return client.Do(req)
}

//POST  method
func HttpPost(url string, requestBody string, params map[string]string, headers map[string]string) (*http.Response, error) {
	var (
		req *http.Request
		err error
	)
	paramStr := ""
	for k, v := range params {
		paramStr = paramStr + k + "=" + v + "&"
	}
	req, err = http.NewRequest("POST", url, strings.NewReader(paramStr))
	if err != nil {
		return nil, errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	// 拼接header参数
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	client := &http.Client{}

	return client.Do(req)
}

func HttpResponseBody(res *http.Response) []byte {
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
