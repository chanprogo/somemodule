package httpclient

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

var HttpClient *http.Client

// createHTTPClient for connection re-use
func CreateHTTPClient(interval int, proxyUrl string) *http.Client {

	if proxyUrl == "" {

		client := &http.Client{

			Transport: &http.Transport{

				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,

				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 50,
				IdleConnTimeout:     time.Duration(60) * time.Second,
			},
			
			Timeout: time.Duration(interval) * time.Second,
		}
		return client
	}

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}

	client := &http.Client{

		Transport: &http.Transport{
			Proxy: proxy,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 50,
			IdleConnTimeout:     time.Duration(60) * time.Second,
		},
		Timeout: time.Duration(interval) * time.Second,
	}
	return client

}
