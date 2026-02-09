dev-crm-api:
	set -a; . ./.env; set +a; go run ./cmd/crm run

dev-project-page-api:
	set -a; . ./.env.project-page; set +a; go run ./cmd/project-page

dev-binder:
	set -a; . ./.env; set +a; go run ./cmd/binder run

dev-crm:
	cd web/crm/ && npm run dev

dev-project-page:
	cd web/project-page/ && npm run dev

gen:
	make gen-crm-api
	make gen-project-page-api

gen-crm-api:
	go generate ./api/crm/...

gen-project-page-api:
	go generate ./api/project-page/...

.PHONY: gen-wire-crm
gen-wire-crm:
	$(call _info,"Generating DI for crm ...")
	@go generate ./cmd/crm

test-api:
	set -a; . ./.env.test; set +a; go test ./... --tags=test,unit

test-api-integration:
	set -a; . ./.env.test; set +a; go test --tags=test,integration ./...
