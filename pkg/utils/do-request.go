package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func (u *Utils) DoRequest(url, httpMethod, path string, data interface{}) ([]byte, error) {
	apiUrl := url + path

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(httpMethod, apiUrl, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	u.addRequestHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (u *Utils) addRequestHeader(req *http.Request) *http.Request {
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	return req
}
