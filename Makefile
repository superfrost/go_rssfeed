# Load env variales from .env if it exist
ifneq (,$(wildcard .env))
  ENV := $(PWD)/.env
  include $(ENV)
else
  $(shell cp ./.env.example ./.env)
  ENV := $(PWD)/.env
  include $(ENV)
endif

APP_NAME=rssagregate
BIN_OUTPUT_DIR=dist

.PHONY: clean
clean: ## Remove temporary files
	go clean -i all
	@rm -rf ${BIN_OUTPUT_DIR}

.PHONY: tidy
tidy: ## Run go mod tidy
	@go mod tidy

.PHONY: vendor
vendor: ## Run go mod vendor
	@go mod vendor

.PHONY: sqlc
sqlc: ## Run sqlc generate
	@sqlc generate

.PHONY: dev
dev: ## Run dev version
	go run ./src/*.go

.PHONY: build
build: clean tidy vendor sqlc ## Build app
	go build -o ./${BIN_OUTPUT_DIR}/${APP_NAME} ./src

.PHONY: run
run: build ## Start App
	./${BIN_OUTPUT_DIR}/${APP_NAME}

.PHONY: doc-infra
doc-infra: ##(local use only) Run infrastructure in docker
	@mkdir -p -m 777 ./_data/data 
	@mkdir -p -m 777 ./_data/pgadmin
	@docker compose -f ./docker-compose.local.yml --env-file ./.env up -d

.PHONY: doc-infra-stop
doc-infra-stop: ##(local use only) Stop infrastructure in docker
	@docker compose -f ./docker-compose.local.yml --env-file ./.env stop

.PHONY: mg-up
mg-up: ## Run goose migrate up
	@goose --dir ./sql/schema postgres "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}" up

.PHONY: mg-down
mg-down: ## Run goose migrate down
	@goose --dir ./sql/schema postgres "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}" down

.PHONY: mg-status
mg-status: ## Run goose migrate status
	@goose --dir ./sql/schema postgres "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}" status
