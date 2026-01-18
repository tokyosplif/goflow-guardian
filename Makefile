.PHONY: run build test docker-up docker-down

run:
	go run cmd/main.go

build:
	go build -o bin/guardian cmd/main.go

test:
	go test -v ./...

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down