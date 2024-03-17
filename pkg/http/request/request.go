package request

import "net/http"

func ParseBearerToken(req *http.Request) string {
	token := req.Header.Get("Authorization")
	if token != "" && len(token) > 7 && token[:7] == "Bearer " {
		return token[7:]
	}
	return ""
}
