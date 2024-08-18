all:
	@go build -o bin/article-collector cmd/main.go

run: all
	@./bin/article-collector

save: all
	@./bin/article-collector save
