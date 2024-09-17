include .env
export

dev:
	env $(cat .env | xargs) go run ./cmd/apartomat run

dev-public-api:
	env $(cat .env.public-api | xargs) go run ./cmd/public-api

gen:
	make gen-crm-api
	make gen-public-api

gen-crm-api:
	go generate ./api/crm/...

gen-public-site-api:
	go generate ./api/public/...

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
