package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Define test cases
	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid name",
			url:            "/hello/John",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hi, John! How you doing?",
		},
		{
			name:           "Empty name",
			url:            "/hello/ ",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Name cannot be empty",
		},
		{
			name:           "Numeric name",
			url:            "/hello/123",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Name should not contain numbers",
		},
	}

	// Create a new router instance
	r := CreateRouter()

	// Run test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a new HTTP request
			req, err := http.NewRequest("GET", test.url, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create a new HTTP recorder
			rec := httptest.NewRecorder()

			// Serve the HTTP request
			r.ServeHTTP(rec, req)

			// Verify the HTTP response status code
			if rec.Code != test.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", rec.Code, test.expectedStatus)
			}

			// Verify the HTTP response body
			if !strings.Contains(rec.Body.String(), test.expectedBody) {
				t.Errorf("Handler returned unexpected body: got %v want %v", rec.Body.String(), test.expectedBody)
			}
		})
	}
}
