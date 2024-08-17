package main

import (
	"github.com/Arey125/article-collector/internal/article"
	_ "github.com/joho/godotenv/autoload"
)

var flyIo = article.Source{
	Url:                 "https://fly.io",
	Domain:              "fly.io",
	ArticleListUrl:      "/blog",
	ArticleListSelector: "article",
	NameSelector:        "h1",
	LinkSelector:        "a",
	ArticleMdSelector:   "article",
}

var goByExample = article.Source{
	Url:                 "https://gobyexample.com/",
	Domain:              "gobyexample.com",
	ArticleListUrl:      "",
	ArticleListSelector: "li",
	NameSelector:        "a",
	LinkSelector:        "a",
	ArticleMdSelector:   "body",
}

var sources = []article.Source{goByExample, flyIo}

func main() {
	for _, source := range sources {
		article.SaveAllArticlesFromSource(source)
	}
}
