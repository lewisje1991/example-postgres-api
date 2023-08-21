.PHONY: up
up:
	@echo "Running..."
	@go run cmd/api/*.go

.PHONY: tools
tools:
	@brew install ariga/tap/atlas@0.13.1
	@brew install sqlc@1.20.0

.PHONY: sqlc
sqlc:
	@sqlc generate

.PHONY: migrate-local-db
migrate-local-db:
	@echo "Migrating database..."
	@atlas schema apply \
		--url "sqlite://file.db" \
		--to "file://schema.sql" \
		--dev-url="sqlite://file.db"