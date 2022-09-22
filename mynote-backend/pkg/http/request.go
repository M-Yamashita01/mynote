package http

import (
	"net/http"
	"strings"
)

func GetBearerTokenFromHeader(request *http.Request) string {
	header := request.Header
	bearToken := header["Authorization"]
	splitBearToken := strings.Split(bearToken[0], " ")
	return splitBearToken[1]
}

func GetRequest(client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, err
	}

	return resp, nil
}
