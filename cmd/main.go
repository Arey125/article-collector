package main

import (
	"os"

	"github.com/Arey125/article-collector/internal/article"
	"github.com/Arey125/article-collector/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	args := os.Args[1:]

    if len(args) == 0 {
        s := server.NewServer()
        s.ListenAndServe()
    }

	if len(args) == 1 && args[0] == "save" {
		for _, source := range article.Sources {
			article.SaveAllArticlesFromSource(source)
		}
        return;
	}
}
