package main

import (
	"fmt"
	fzf "github.com/ktr0731/go-fuzzyfinder"
	"os"
    "github.com/Arey125/article-collector/internal/article"
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

func main() {
	articles, err := goByExample.GetArticleList()
	if err != nil {
		panic(err)
	}

	selectedInd, err := fzf.Find(articles, func(i int) string {
		return articles[i].Name
	})
	if err != nil {
		panic(err)
	}
	selectedArticle := articles[selectedInd]
	articleMd, err := selectedArticle.GetMd()
	if err != nil {
		panic(err)
	}

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println(articleMd)
		return
	}
	outputFile, err := os.Create(args[0])
	if err != nil {
		panic(err)
	}

	outputFile.Write([]byte(articleMd))
}
