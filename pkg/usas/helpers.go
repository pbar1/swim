package usas

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// TODO: set `Referer` at request site
var headers = map[string]string{
	"RequestVerificationToken": "2YQ5Gs9TY6WCD7LrWF6IqjE5DidlcPMS5R5lO5sKW2aD6nG-oHVJLawtGjNkR7UWuFcNQUSvp-Q0D7KzjBPFiLzsbfQ1:I4ZUkMDpZxelWutmPFyKmcpFJv5GIX38MPFOvJllshQBuNwRgwt-ew86qJr9mKnoBPijxw06DmyB3PBvG8oq7fjtZf81",
	"Origin":                   "https://www.usaswimming.org",
	"User-Agent":               "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
	"Content-Type":             "application/x-www-form-urlencoded",
	"Accept":                   "*/*",
	"X-Requested-With":         "XMLHttpRequest",
	"Accept-Encoding":          "gzip, deflate, br",
	"Accept-Language":          "en-US,en;q=0.9",
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
