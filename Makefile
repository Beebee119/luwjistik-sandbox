tidy:
	go mod tidy

start:
	go run ./bin/web

build:
	go build ./bin/web

test:
	go test -v ./...
