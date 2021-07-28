package utils

import (
	"net/http"
)

func ParseBody(request *http.Request) (string, error) {
	buf := make([]byte, 1024)
	num, _ := request.Body.Read(buf)
	str := string(buf[0:num])
	return str, nil
}
