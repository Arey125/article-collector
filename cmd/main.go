package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Arey125/article-collector/internal/article"
	"github.com/Arey125/article-collector/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	args := os.Args[1:]

    if len(args) == 0 {
        s := server.NewServer()
        port, _ := strconv.Atoi(os.Getenv("PORT"))
        fmt.Printf("listening on http://127.0.0.1:%d\n", port)
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
