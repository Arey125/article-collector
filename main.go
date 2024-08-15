package main

import (
	"fmt"
	"net/http"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

type Article struct {
    name string
    link string
}

func getArticleMd(url string, selector string) (string, error) {
    fmt.Println(url)
    res, err := http.Get(url)
    if err != nil {
        return "", err
    }

    doc, err := goquery.NewDocumentFromResponse(res)
    if err != nil {
        return "", err
    }

    articleElement := doc.Find(selector)
    converter := md.NewConverter("", true, nil)

    return converter.Convert(articleElement), nil
}

func getArticleList (url string, articleSelector string, nameSelector string, linkSelector string) ([]Article, error) {
    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    doc, err := goquery.NewDocumentFromResponse(res)
    if err != nil {
        return nil, err
    }

    articleSelect := doc.Find(articleSelector)
    articles := make([]Article, len(articleSelect.Nodes), len(articleSelect.Nodes))
    for i := range articleSelect.Nodes {
        articleNode := articleSelect.Eq(i)
        articles[i] = Article{
            name: strings.TrimSpace(articleNode.Find(nameSelector).Text()),
            link: articleNode.Find(linkSelector).AttrOr("href", ""),
        }
    }

    return articles, nil
}

func main() {
    url := "https://fly.io"

    articles, err := getArticleList(url + "/blog", "article", "h1", "a")
    if (err != nil) {
        panic(err)
    }

    fmt.Println(getArticleMd(url + articles[0].link, "article > section"))
}
