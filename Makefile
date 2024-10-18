include .env
export

dev:
	env $(cat .env | xargs) go run ./cmd/crm run

dev-project-page:
	env $(cat .env.project-page | xargs) go run ./cmd/project-page

gen:
	make gen-crm-api
	make project-page-api

gen-crm-api:
	go generate ./api/crm/...

gen-project-page-api:
	go generate ./api/project-page/...

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
