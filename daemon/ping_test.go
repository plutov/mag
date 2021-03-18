package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingTarget(t *testing.T) {
	testServer500 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer500.Close()

	testServer200 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer testServer200.Close()

	var tests = []struct {
		endpoint    string
		expectError bool
	}{
		{testServer500.URL, true},
		{testServer200.URL, false},
	}

	for _, tt := range tests {
		t.Run(tt.endpoint, func(t *testing.T) {
			err := PingTarget(ConfigEntry{
				Endpoint:         tt.endpoint,
				ExpectStatusCode: 200,
				Method:           "GET",
			})

			gotError := err != nil

			if gotError != tt.expectError {
				t.Errorf("expect error: %t, got error: %t, error: %v", tt.expectError, gotError, err)
			}
		})
	}
}
