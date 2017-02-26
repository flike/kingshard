all: build

build: kingshard
goyacc:
	go build -o ./bin/goyacc ./vendor/golang.org/x/tools/cmd/goyacc
kingshard: goyacc
	./bin/goyacc -o ./sqlparser/sql.go ./sqlparser/sql.y
	gofmt -w ./sqlparser/sql.go
	@bash genver.sh
	go build -o ./bin/kingshard ./cmd/kingshard
clean:
	@rm -rf bin
	@rm -f ./sqlparser/y.output ./sqlparser/sql.go

test:
	go test ./go/... -race
