.PHONY: run proto sqlc-generate \
        migrate-up migrate-down migrate-up-steps migrate-down-steps migrate-goto \
        migrate-version migrate-force \
        migrate-validate migrate-up-and-validate migrate-up-steps-and-validate

DB_URL := postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable
MIGRATIONS_DIR := ./database/migrations

MIGRATE := docker compose --profile tools run --rm migrate

PSQL := docker compose exec -T postgres psql "$(DB_URL)"

run:
	go run cmd/server/main.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/server/v1/user.proto

sqlc-generate:
	sqlc generate

migrate-validate:
	@set -euo pipefail; \
	echo "Running validation scripts..."; \
	for f in $(MIGRATIONS_DIR)/*_*.validate.sql; do \
		[ -e "$$f" ] || { echo "No validate files, skip"; break; } ; \
		echo "  > $$f"; \
		out="$$( cat "$$f" | $(PSQL) -v ON_ERROR_STOP=1 -qAt | tail -n1 )"; \
		if [ "$$out" != "OK" ]; then \
			echo "Validation failed for $$f (got '$$out')"; \
			exit 1; \
		fi; \
	done; \
	echo "Validation OK"

migrate-up:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migrate-down:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

migrate-up-steps:
	@read -p "Enter number of steps to migrate up: " steps; \
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up $$steps

migrate-down-steps:
	@read -p "Enter number of steps to migrate down: " steps; \
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down $$steps

migrate-goto:
	@read -p "Enter target version (e.g., 1, 2, 3): " version; \
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" goto $$version

migrate-version:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

migrate-force:
	@read -p "Enter version to force (e.g., 1): " version; \
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $$version

migrate-up-and-validate:
	@set -e; \
	echo "== up 1 step =="; \
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up 1; \
	echo "== validate =="; \
	if $(MAKE) --no-print-directory migrate-validate; then \
		echo "Migrate + Validate OK"; \
	else \
		echo "Validate failed, rolling back 1 step..."; \
		$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1 || true; \
		exit 1; \
	fi

migrate-up-steps-and-validate:
	@set -e; \
	read -p "Enter number of steps to migrate up (validated each step): " steps; \
	i=0; \
	while [ $$i -lt $$steps ]; do \
		$(MAKE) --no-print-directory migrate-up-and-validate; \
		i=$$((i+1)); \
	done
