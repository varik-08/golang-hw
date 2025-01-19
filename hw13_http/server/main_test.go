package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetRequest(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		headers      map[string]string
		queryParams  string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid GET request",
			method:       http.MethodGet,
			headers:      nil,
			queryParams:  "",
			expectedCode: http.StatusOK,
			expectedBody: "New request:\r\nMethod: GET",
		},
		{
			name:         "Invalid method (POST)",
			method:       http.MethodPost,
			headers:      nil,
			queryParams:  "",
			expectedCode: http.StatusBadRequest,
			expectedBody: "Method not allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/?"+tt.queryParams, nil)

			recorder := httptest.NewRecorder()

			getRequest(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedCode, res.StatusCode)

			body := recorder.Body.String()
			assert.Contains(t, body, tt.expectedBody)
		})
	}
}

func Test_PostRequest(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		headers      map[string]string
		queryParams  string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid Post request",
			method:       http.MethodPost,
			headers:      nil,
			queryParams:  "",
			expectedCode: http.StatusOK,
			expectedBody: "New request:\r\nMethod: POST",
		},
		{
			name:         "Invalid method (GET)",
			method:       http.MethodGet,
			headers:      nil,
			queryParams:  "",
			expectedCode: http.StatusBadRequest,
			expectedBody: "Method not allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/?"+tt.queryParams, nil)

			recorder := httptest.NewRecorder()

			postRequest(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedCode, res.StatusCode)

			body := recorder.Body.String()
			assert.Contains(t, body, tt.expectedBody)
		})
	}
}
