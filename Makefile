all:
	@go build -o bin/article-collector cmd/main.go

run: all
	./bin/article-collector
