.PHONY: run-docker-local
run-docker-local:
	docker compose --env-file .env -f docker-compose.yml up --build --remove-orphans
