NAME=calculator

test:
	go test

format:
	gofmt -d .

run:
	go run cmd/calculator/main.go

build:
	go build -o bin/$(NAME) ./cmd/calculator