# build
FROM golang:1.14 as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/plutov/mag/daemon

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o daemon .

# binary only
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/src/github.com/plutov/mag/daemon .

CMD ["./daemon"]