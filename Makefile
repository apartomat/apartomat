include .env
export

dev:
	SERVER_ADDR=localhost:8010 go run ./cmd/apartomat run

gen:
	go generate ./...

build:
	GOOS=linux GOARCH=amd64 go build -o bin/apartomat-lunux-amd64 ./cmd/apartomat
