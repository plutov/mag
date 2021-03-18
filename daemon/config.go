package main

import (
	"encoding/json"
	"os"
)

// ConfigEntry .
type ConfigEntry struct {
	Endpoint         string `json:"endpoint"`
	Method           string `json:"method"`
	FrequencySeconds int    `json:"frequencySeconds"`
	ExpectStatusCode int    `json:"expectStatusCode"`
	Timeout          int    `json:"timeout"`
	FailureThreshold int    `json:"failureThreshold"`
	FailuresCounter  int    `json:"-"`
}

// ReadConfigFile .
func ReadConfigFile() ([]ConfigEntry, error) {
	file, err := os.Open(os.Getenv("TARGETS_CONFIG_FILE"))
	if err != nil {
		return nil, err
	}

	res := []ConfigEntry{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}
