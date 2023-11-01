run:
	@go build -o bin/gobank

build:
	@./bin/gobank

test:
	@go test -v ./...
