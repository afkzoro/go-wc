.PHONY: build test clean

build:
	go build -o bin/go-wc cmd/wc/main.go

test:
	go test ./...

clean:
	rm -rf bin/