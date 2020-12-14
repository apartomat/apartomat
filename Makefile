include .env
export

dev:
	go run ./cmd/apartomat run

gen:
	go generate ./...