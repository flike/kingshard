#!/bin/bash
export GOBIN=/usr/local/go/bin
export GOARCH=amd64
export GOROOT=/usr/local/go
export GOOS=linux

source ./dev.sh

go tool yacc -o ./sqlparser/sql.go ./sqlparser/sql.y
gofmt -w ./sqlparser/sql.go
source  genver.sh
go build -o ./bin/ucdbproxy ./cmd/kingshard


