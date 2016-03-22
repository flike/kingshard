all: build

build: kingshard

kingshard:
	@bash genver.sh
	go build -o ./bin/kingshard ./cmd/kingshard
clean:
	@rm -rf bin

test:
	go test ./go/... -race
