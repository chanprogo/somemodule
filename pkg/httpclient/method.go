package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func PostByteJson(baseURL string, url string, data []byte, setter ...func(*http.Request)) (body []byte, header http.Header, statusCode int) {
	return reqJsonData(baseURL, "POST", url, data, setter...)
}

func reqJson(baseURL string, method string, url string, data []byte, setters ...func(*http.Request)) (body []byte, header http.Header, statusCode int) {

	url = baseURL + url
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	if method == "POST" || method == "PUT" || method == "DELETE" {
		request.Header.Set("Content-Type", "application/json;charset=utf8;")
	}

	return reqInner(request, setters...)
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

func reqInner(request *http.Request, setters ...func(*http.Request)) (body []byte, header http.Header, statusCode int) {
	var (
		err error
		res *http.Response
	)

	for _, setter := range setters {
		setter(request)
	}

	// log request

	res, err = HttpClient.Do(request)
	// request.URL.String()
	if err != nil {
		// log err.Error()
		return
	}

	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		// log err.Error()
		return
	}

	// log string(body)

	header = res.Header
	statusCode = res.StatusCode

	return
}
