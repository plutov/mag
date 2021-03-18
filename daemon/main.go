package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)

	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
}

func main() {
	config, err := ReadConfigFile()
	if err != nil {
		log.WithError(err).Error("unable to read config file")
		os.Exit(1)
	}

	// Start timers to ping each target concurrently in a go routine
	for _, target := range config {
		logCtx := log.WithField("endpoint", target.Endpoint)
		logCtx.Info("target is registered")

		go func(t ConfigEntry, l *log.Entry) {
			ticker := time.NewTicker(time.Second * time.Duration(t.FrequencySeconds))
			defer ticker.Stop()

			for range ticker.C {
				err := PingTarget(t)
				if err != nil {
					l.WithError(err).Error("healthcheck failed")
					t.FailuresCounter++
					if t.FailuresCounter >= t.FailureThreshold {
						l.Error("endpoint is not responding, target is de-registered")
						ticker.Stop()
					}
				} else {
					l.Debug("ok")
					t.FailuresCounter = 0
				}
			}
		}(target, logCtx)
	}

	// @TODO: graceful shutdown
	select {} // block forever, since it's a daemon and should run unless we stop it
}
