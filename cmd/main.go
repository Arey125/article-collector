package main

import (
	"fmt"
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
        return;
    }

	if len(args) == 1 && args[0] == "save" {
        fmt.Println("saving articles...")
		for _, source := range article.Sources {
			article.SaveAllArticlesFromSource(source)
		}
        fmt.Println("done")
        return;
	}
}
