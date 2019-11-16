COMMIT_HASH=$(shell git rev-parse --short HEAD || echo "GitNotFound")
BUILD_DATE=$(shell date '+%Y-%m-%d %H:%M:%S')

all: build

build: kingshard
goyacc:
	go get -u golang.org/x/tools/cmd/goyacc
	${GOPATH}/bin/goyacc -o ./sqlparser/sql.go ./sqlparser/sql.y
	gofmt -w ./sqlparser/sql.go
kingshard:
	go build -mod=vendor -ldflags "-X \"main.BuildVersion=${COMMIT_HASH}\" -X \"main.BuildDate=$(BUILD_DATE)\"" -o ./bin/kingshard ./cmd/kingshard
clean:
	@rm -rf bin
test:
	go test ./go/... -race