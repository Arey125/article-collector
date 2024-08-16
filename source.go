package main

import (
	"net/http"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

type Source struct {
	url    string
	domain string

	articleListUrl      string
	articleListSelector string
	nameSelector        string
	linkSelector        string

	articleMdSelector string
}

type Article struct {
	name   string
	link   string
	source *Source
}

func (source Source) getArticleList() ([]Article, error) {
	res, err := http.Get(source.url + source.articleListUrl)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	articleSelect := doc.Find(source.articleListSelector)
	articles := make([]Article, len(articleSelect.Nodes), len(articleSelect.Nodes))
	for i := range articleSelect.Nodes {
		articleNode := articleSelect.Eq(i)
		articles[i] = Article{
			name:   strings.TrimSpace(articleNode.Find(source.nameSelector).Text()),
			link:   articleNode.Find(source.linkSelector).AttrOr("href", ""),
			source: &source,
		}
	}

	return articles, nil
}

func (article Article) getMd() (string, error) {
	res, err := http.Get(article.source.url + article.link)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	articleElement := doc.Find(article.source.articleMdSelector)
	converter := md.NewConverter(article.source.domain, true, nil)

	return converter.Convert(articleElement), nil
}
