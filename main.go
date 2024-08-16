package main

import (
	"fmt"
	fzf "github.com/ktr0731/go-fuzzyfinder"
	"os"
)

var flyIo = Source{
	url:                 "https://fly.io",
	domain:              "fly.io",
	articleListUrl:      "/blog",
	articleListSelector: "article",
	nameSelector:        "h1",
	linkSelector:        "a",
	articleMdSelector:   "article",
}

var goByExample = Source{
	url:                 "https://gobyexample.com/",
	domain:              "gobyexample.com",
	articleListUrl:      "",
	articleListSelector: "li",
	nameSelector:        "a",
	linkSelector:        "a",
	articleMdSelector:   "body",
}

func main() {
	articles, err := goByExample.getArticleList()
	if err != nil {
		panic(err)
	}

	selectedInd, err := fzf.Find(articles, func(i int) string {
		return articles[i].name
	})
	if err != nil {
		panic(err)
	}
	selectedArticle := articles[selectedInd]
	articleMd, err := selectedArticle.getMd()
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
