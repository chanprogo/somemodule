package httpclient

import (
	"io/ioutil"
	"net/http"
)

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
