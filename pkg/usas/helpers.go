package usas

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// TODO: set `Referer` at request site
var headers = map[string]string{
	"Accept":                   "*/*",
	"Accept-Encoding":          "gzip, deflate, br",
	"Accept-Language":          "en-US,en;q=0.5",
	"Connection":               "keep-alive",
	"Content-Type":             "application/x-www-form-urlencoded; charset=UTF-8",
	"Host":                     "www.usaswimming.org",
	"Referer":                  "https://www.usaswimming.org/times/event-rank-search",
	"RequestVerificationToken": "OXJ0vlRBcNyzj7NmZamU9_oQX7cd4il5MEnM5--k59vf4hkIIe-YChdRy7nD4ZkbUiNlxiZIbVJplYIU1jthpkdmSMM1:ceLtdeq6KVGe55VmTqRbttJR8mxfTaCO6_ynPipXtARf1uhapildOZnsFMW2gErLJwfTu_PZiCQDxMLFnC_zrgaP6VE1",
	"User-Agent":               "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
	"X-Requested-With":         "XMLHttpRequest",
}

type ErrReqVerTokenNotFound struct {
	Resp http.Response
}

func (e ErrReqVerTokenNotFound) Error() string {
	return fmt.Sprintf("request verification token not in http response: %v", e.Resp)
}

func getReqVerToken() (*string, error) {
	resp, err := http.Get("https://www.usaswimming.org/")
	if err != nil {
		return nil, err
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "__RequestVerificationToken" {
			return &cookie.Value, nil
		}
	}

	return nil, ErrReqVerTokenNotFound{*resp}
}

func makePost(reqURL string, formData url.Values) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", reqURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
