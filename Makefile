include .env
export

dev:
	env $(cat .env | xargs) go run ./cmd/apartomat run

gen:
	go generate ./...

test:
	env $(cat .env | xargs) go test ./... --tags=test,unit

test-integration:
	env $(cat .env | xargs) go test --tags=test,integration ./...

build:
	GOOS=linux GOARCH=amd64 go build -o bin/apartomat-lunux-amd64 ./cmd/apartomat

docker:
	GOOS=linux GOARCH=amd64 go build -o bin/apartomat-lunux-amd64 ./cmd/apartomat
	GOOS=linux GOARCH=amd64 go build -o bin/migration-lunux-amd64 ./cmd/migration
	docker build . -t apartomat

docker-test:
	docker build . -t test --file Dockerfile-test