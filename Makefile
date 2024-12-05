ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DB_CONN=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST)/$(DB_NAME)?sslmode=disable

status:
	@echo "Running status command"
	@echo "DB_CONN=$(DB_CONN)"
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_CONN) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) status

reset:
	@echo "Running reset command"
	@echo "DB_CONN=$(DB_CONN)"
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_CONN) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) reset

down:
	@echo "Running down command"
	@echo "DB_CONN=$(DB_CONN)"
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_CONN) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) down

up:
	@echo "Running up command"
	@echo "DB_CONN=$(DB_CONN)"
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_CONN) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) up


.PHONY: run-docker-local
run-docker-local:
	docker compose --env-file .env -f docker-compose.yml up --build --remove-orphans
