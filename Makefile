include .env

all:
	@templ generate
	@go build -o bin/article-collector cmd/main.go

run: all
	@./bin/article-collector

save: all
	@./bin/article-collector save

mkmgr:
	@./scripts/make_migration.sh

sqlite:
	@sqlite3 $(DB)
