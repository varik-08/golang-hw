package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetRequest(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello world"}`))
	}))
	defer mockServer.Close()

	parsedURL, _ := url.Parse(mockServer.URL)

	hostPort := parsedURL.Host

	err := getRequest("http", hostPort, "mock")

	require.Nil(t, err)
}

func Test_PostRequest(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello world"}`))
	}))
	defer mockServer.Close()
	parsedURL, _ := url.Parse(mockServer.URL)

	hostPort := parsedURL.Host

	err := postRequest("http", hostPort, "mock", `{"key": "value"}`)

	require.Nil(t, err)
}
