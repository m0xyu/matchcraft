.PHONY: test lint run-example

test:
	go test -v -cover ./...

lint:
	golangci-lint run

run-example:
	go run examples/coffee/main.go

cover:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out