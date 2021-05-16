package httpclient

import (
	"bytes"
	"net/http"
)

func PostByteJson(baseURL string, url string, data []byte, setter ...func(*http.Request)) (body []byte, header http.Header, statusCode int) {
	return reqJsonData(baseURL, "POST", url, data, setter...)
}

func reqJsonData(baseURL string, method string, url string, data []byte, setters ...func(*http.Request)) (body []byte, header http.Header, statusCode int) {
	url = baseURL + url
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	if method == "POST" || method == "PUT" || method == "DELETE" {
		request.Header.Set("Content-Type", "multipart/form-data;charset=utf8;")
	}
	return reqInner(request, setters...)
}
