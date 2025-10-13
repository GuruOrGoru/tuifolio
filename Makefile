.PHONY: all build clean test
.SILENT: build clean test run

all: build

build:
	go build -o tuifolio ./cmd/server

clean:
	rm -f tuifolio

test:
	go test ./...

run: build
	./tuifolio
