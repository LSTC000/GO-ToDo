GO=go
GO_MOD=$(GO) mod

APP_CMD_NAME=main.go
APP_CMD_PATH=cmd/app
APP_BUILD_NAME=main
APP_BUILD_PATH=build/app

DC=docker compose
DC_LOCAL=$(DC) -f ./deployment/docker-compose.local.yaml


.PHONY: app
app-run:
	$(GO) run ./$(APP_CMD_PATH)/$(APP_CMD_NAME)
app-build:
	$(GO) build -C ./$(APP_CMD_PATH) -o ../../$(APP_BUILD_PATH)/$(APP_BUILD_NAME)
app-start:
	./$(APP_BUILD_PATH)/$(APP_BUILD_NAME)


.PHONY: mod
init:
	$(GO_MOD) init
tidy:
	$(GO_MOD) tidy
vendor:
	$(GO_MOD) vendor
download:
	$(GO_MOD) download


.PHONY: lint
import-run:
	goimports -w -l ./..
fmt-run:
	gofmt -w -l ./..
lint-run:
	golangci-lint run

.PHONY: docs
swag-gen:
	swag init -g ./cmd/app/main.go

.PHONY: docker
docker-local:
	$(DC_LOCAL) up -d --build
docker-local-build:
	$(DC_LOCAL) build
docker-local-up:
	$(DC_LOCAL) up -d
docker-local-stop:
	$(DC_LOCAL) stop
docker-local-start:
	$(DC_LOCAL) start
docker-local-down:
	$(DC_LOCAL) down
docker-local-down-v:
	$(DC_LOCAL) down -v
docker-local-logs:
	$(DC_LOCAL) logs
docker-local-logs-f:
	$(DC_LOCAL) logs -f
