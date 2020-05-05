FROM golang:1.11-alpine

ENV GOPROXY="https://proxy.golang.org"
ENV GO111MODULE="on"
WORKDIR /go/src/github.com/icco/tokei

RUN apk add --no-cache git
COPY . .

RUN go build -v ./cmd/tokei
RUN go build -v ./cmd/inky
