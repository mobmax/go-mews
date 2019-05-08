// +build !integration

package accountingcategories

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

const (
	demoBaseURL = "https://demo.mews.li/api/connector/v1/"
	token       = "C66EF7B239D24632943D115EDE9CB810-EA00F8FD8294692C940F6B5A8F9453D"
)

var (
	mux     *http.ServeMux
//	client  *Service
	server  *httptest.Server
	service *Service
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	baseURL, _ := url.Parse(demoBaseURL)

	service = NewService()
	service.Client.BaseURL = baseURL
	service.Client.AccessToken = token

	// set custom http client
	service.Client.Client = &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	got := r.Method
	if expected != got {
		t.Errorf("Request method = %v, expected %v", got, expected)
	}
}

func testHeader(t *testing.T, r *http.Request, key string, expected string) {
	got := r.Header.Get(key)
	if expected != got {
		t.Errorf("Request header %v = %v, expected %v", key, got, expected)
	}
}
