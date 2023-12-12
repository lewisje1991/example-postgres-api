.PHONY: gen-env-local
gen-env-local:
	@echo "Generating .env local file..."
	@echo "MODE=non-prod" > .env
	@echo "DB_URL=postgres://postgres:postgres@db:5432/code-bookmarks?sslmode=disable" >> .env

.PHONY: up
up: 
	@templ generate
	@docker compose run --build -p 8080:8080 api 

.PHONY: tools
tools:
	@brew install ariga/tap/atlas
	@brew install sqlc
	@brew install orbstack

.PHONY: sqlc
sqlc:
	@sqlc generate

.PHONY: migrate-local
migrate-local:
	@echo "Migrating database..."
	@atlas schema apply --env local


# .PHONY: migrate-dev
# migrate-dev:
# 	@echo "Migrating database..."
# 	@set -eu; \
# 	DB_TOKEN=$$(turso db tokens create code-bookmarks-dev); \
# 	echo "DB token created."; \
# 	atlas schema apply --auto-approve --env dev --to file://schema.sql --dev-url "sqlite://dev?mode=memory"; \
# 	echo "Schema applied."