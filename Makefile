all: build

build: kingshard

kingshard:
	go build -ldflags "-X main.GitHash=`git rev-parse HEAD` -X main.BuildTime=`date '+%Y-%m-%d_%I:%M:%S%p'`" -o ./bin/kingshard ./cmd/kingshard

clean:
	@rm -rf bin

test:
	go test ./go/... -race
