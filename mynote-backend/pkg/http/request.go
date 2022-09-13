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
