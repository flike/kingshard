all: build

build: kingshard

kingshard:
	go build -o bin/kingshard ./cmd/kingshard

clean:
	@rm -rf bin

test:
	go test ./go/... -race
