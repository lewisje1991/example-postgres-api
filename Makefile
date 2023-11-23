.PHONY: gen-env-local
gen-env-local:
	@echo "Generating .env local file..."
	@echo "MODE=non-prod" > .env
	@echo "DB_URL=$$(turso db show code-bookmarks-dev --url)" >> .env
	@echo "DB_TOKEN=$$(turso db tokens create code-bookmarks-dev)" >> .env

.PHONY: up
up: gen-env-local
	@docker compose run -p 8080:8080 api

.PHONY: tools
tools:
	@brew install tursodatabase/tap/turso
	@brew install ariga/tap/atlas
	@brew install sqlc
	@brew install flyctl
	@brew install orbstack

.PHONY: sqlc
sqlc:
	@sqlc generate

.PHONY: migrate-dev
migrate-dev:
	@echo "Migrating database..."
	@DB_TOKEN=$$(turso db tokens create code-bookmarks-dev) && docker compose run migrate

# .PHONY: migrate-dev
# migrate-dev:
# 	@echo "Migrating database..."
# 	@set -eu; \
# 	DB_TOKEN=$$(turso db tokens create code-bookmarks-dev); \
# 	echo "DB token created."; \
# 	atlas schema apply --auto-approve --env dev --to file://schema.sql --dev-url "sqlite://dev?mode=memory"; \
# 	echo "Schema applied."