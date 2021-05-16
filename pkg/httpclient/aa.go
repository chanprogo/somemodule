package httpclient

import (
	"net/http"
)

func PostJson(baseURL string, url string, data []byte, setter ...func(*http.Request)) (body []byte, header http.Header, statusCode int) {
	return reqJson(baseURL, "POST", url, data, setter...)
}

func Get(baseURL string, url string, data []byte, setter ...func(*http.Request)) (body []byte, header http.Header, statusCode int) {
	return reqJson(baseURL, "GET", url, data, setter...)
}
