package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Arey125/article-collector/internal/article"
	"github.com/Arey125/article-collector/internal/models"
	"github.com/Arey125/article-collector/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	args := os.Args[1:]
	db, err := models.InitDb()
	if err != nil {
		panic(err)
	}

	if len(args) == 0 {
		s := server.NewServer(db)
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		fmt.Printf("listening on http://127.0.0.1:%d\n", port)
		s.ListenAndServe()
		return
	}

	if len(args) == 1 && args[0] == "save" {
		fmt.Println("saving articles...")
		for _, source := range article.Sources {
			saveAllArticlesFromSource(*source, db)
		}
		fmt.Println("done")
		return
	}
}

func saveAllArticlesFromSource(source article.Source, db *sql.DB) error {
	articleModel := models.ArticleModel{DB: db}

	articles, err := source.GetArticleListFromHtml()
	if err != nil {
		return err
	}

	for _, article := range articles {
		err := articleModel.InsertOrReplace(&article)
        if err != nil {
            panic(err)
        }

		saved, err := article.SaveToFileIfDoesNotExist()
		if err != nil {
			panic(err)
		}
		if saved {
			fmt.Printf("\"%s\" from %s is saved\n", article.Name, article.Source.Domain)
			time.Sleep(400 * time.Millisecond)
		}
	}

	return nil
}
