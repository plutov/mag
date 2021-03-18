package main

import (
	"fmt"
	"net/http"
	"time"
)

// PingTarget .
func PingTarget(target ConfigEntry) error {
	httpClient := http.Client{
		Timeout: time.Second * time.Duration(target.Timeout),
	}

	req, err := http.NewRequest(target.Method, target.Endpoint, nil)
	if err != nil {
		return err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != target.ExpectStatusCode {
		return fmt.Errorf("expecting status code %d, got %d", target.ExpectStatusCode, res.StatusCode)
	}

	return nil
}
