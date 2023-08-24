.PHONY: up
up:
	@echo "Running..."
	@go run cmd/api/*.go

.PHONY: tools
tools:
	@brew install ariga/tap/atlas
	@brew install sqlc
	@brew install flyctl

.PHONY: sqlc
sqlc:
	@sqlc generate

.PHONY: show-db-info
show-db-info:
	@turso db show code-bookmarks

.PHONY: auth-db
auth-db:
	@export TURSO_TOKEN=$(turso db tokens create code-bookmarks)

.PHONY: connect-db
connect-db:
	@turso db shell code-bookmarks

.PHONY: migrate-local-db
migrate-local-db:
	@echo "Migrating database..."
	@atlas schema apply \
		--url "sqlite://file.db" \
		--to "file://schema.sql" \
		--dev-url="sqlite://file.db"

.PHONY: migrate-db
migrate-db:
	@echo "Migrating database..."
	@atlas schema apply --env turso --to file://schema.sql --dev-url "sqlite://dev?mode=memory"

