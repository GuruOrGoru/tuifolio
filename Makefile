.PHONY: all build clean test

all: build

build:
	go build -o tuifolio ./cmd/server

clean:
	rm -f tuifolio

test:
	go test ./...

run:
	go run cmd/server/main.go
