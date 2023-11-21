.PHONY: gen-env-local
gen-env-local:
	@echo "Generating .env local file..."
	@echo "MODE=non-prod" > .env
	@echo "DB_URL=libsql://db:5001" >> .env

.PHONY: gen-env-prod
gen-env-prod:
	@echo "Generating .env prod file..."
	@echo "MODE=prod" > .env
	@echo "DB_URL=$$(turso db show code-bookmarks --url)" >> .env
	@echo "DB_TOKEN=$$(turso db tokens create code-bookmarks)" >> .env

.PHONY: up
up:
	@go run cmd/api/*.go

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

.PHONY: inspect-db
inspect-db:
	@turso db show code-bookmarks

.PHONY: connect-db
connect-db:
	@turso db shell code-bookmarks


.PHONY: migrate-local
migrate-local:
	@echo "Migrating database..."
	@atlas schema apply --env local --to file://schema.sql --dev-url "sqlite://dev?mode=memory"


.PHONY: migrate-prod
migrate-prod:
	@echo "Migrating database..."
	@export DB_TOKEN=$$(turso db tokens create code-bookmarks) && \
		atlas schema apply --env turso --to file://schema.sql --dev-url "sqlite://dev?mode=memory"

