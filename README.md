# mag

A daemon to check the status of REST API. Using the set of defined targets it checks concurrently each endpoint and logs the results when debugging is enabled.

Features:
- check status code
- configurable HTTP methods, multiple endpoints
- support timeouts
- failure threshold - if there are too many failures - stop checking the endpoint

Example log output:
```
docker-compose logs -f 
time="2021-03-18T18:19:50Z" level=info msg="target is registered" endpoint="http://magnificent:12345"
time="2021-03-18T18:19:50Z" level=info msg="target is registered" endpoint="http://httpbin.org/status/500"

time="2021-03-18T18:19:55Z" level=debug msg=ok endpoint="http://magnificent:12345"
time="2021-03-18T18:19:55Z" level=error msg="healthcheck failed" endpoint="http://httpbin.org/status/500" error="expecting status code 200, got 500"

...

time="2021-03-18T18:20:15Z" level=error msg="endpoint is not responding, target is de-registered" endpoint="http://httpbin.org/status/500

...

time="2021-03-18T18:21:15Z" level=error msg="healthcheck failed" endpoint="http://magnificent:12345" error="expecting status code 200, got 500"
```

## Requirements

- Docker with or without Docker Compose

## Configuration

Daemon can be configured using environment variables.

- `TARGETS_CONFIG_FILE` - path to the file with targets config
- `LOG_LEVEL` - INFO/DEBUG/...

targets.json example:

```json
[[
  {
    "endpoint": "http://magnificent:12345",
    "method": "GET",
    "expectStatusCode": 200,
    "frequencySeconds": 5,
    "timeout": 30,
    "failureThreshold": 10
  },
  {
    "endpoint": "http://httpbin.org/status/500",
    "method": "GET",
    "expectStatusCode": 200,
    "frequencySeconds": 5,
    "timeout": 30,
    "failureThreshold": 5
  }
]
```

## Run with Docker

```
# build and run in the background
docker-compose up -d --build

# check logs
docker-compose logs -f daemon
```

## Run unit tests

```
cd ./daemon
go test -v ./...

=== RUN   TestPingTarget
=== RUN   TestPingTarget/http://127.0.0.1:56000
=== RUN   TestPingTarget/http://127.0.0.1:56001
--- PASS: TestPingTarget (0.00s)
    --- PASS: TestPingTarget/http://127.0.0.1:56000 (0.00s)
    --- PASS: TestPingTarget/http://127.0.0.1:56001 (0.00s)
PASS
ok  	daemon	0.292s
```

## TODO

- Extend target parameters, allow more configuration, such as Headers, etc.
- Store results into time-series DB in order to be able to visualize it later
- If there are many targets, setup concurrency policy to prevent a lot of requests at the same time (if network doesn't allow)

## Changelog

- 30 mins: setup docker environment for daemon and magnificent server
- 1 hour: read config file, do healthchecks in the background independently, report as logs
- 1 hour 30 mins: manual testing, adding failureThreshold config paramter to tell that endpoint fails for N attempts
- 2 hours: basic unit tests for PingTarget func