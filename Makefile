.PHONY: gen-env-local
gen-env-local:
	@echo "Generating .env local file..."
	@echo "MODE=non-prod" > .env
	@echo "DB_URL=postgres://postgres:postgres@db:5432/code-bookmarks?sslmode=disable" >> .env

.PHONY: up
up:
	@docker compose run --build -p 8080:8080 api 

.PHONY: sqlc
sqlc:
	@sqlc generate

.PHONY: migrate-local
migrate-local:
	@echo "Migrating database..."
	@atlas schema apply --env local --var password=$(password)

.PHONY: tools
tools:
	@brew install ariga/tap/atlas
	@brew install sqlc
	@brew install orbstack
	@brew install bruno