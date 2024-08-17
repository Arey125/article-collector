package article

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Source struct {
	Url    string
	Domain string

	ArticleListUrl      string
	ArticleListSelector string
	NameSelector        string
	LinkSelector        string

	ArticleMdSelector string
}

func (source Source) GetArticleList() ([]Article, error) {
	res, err := http.Get(source.Url + source.ArticleListUrl)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	articleSelect := doc.Find(source.ArticleListSelector)
	articles := make([]Article, len(articleSelect.Nodes), len(articleSelect.Nodes))
	for i := range articleSelect.Nodes {
		articleNode := articleSelect.Eq(i)
		articles[i] = Article{
			Name:   strings.TrimSpace(articleNode.Find(source.NameSelector).Text()),
			Link:   articleNode.Find(source.LinkSelector).AttrOr("href", ""),
			Source: &source,
		}
	}

	return articles, nil
}

var flyIo = Source{
	Url:                 "https://fly.io",
	Domain:              "fly.io",
	ArticleListUrl:      "/blog",
	ArticleListSelector: "article",
	NameSelector:        "h1",
	LinkSelector:        "a",
	ArticleMdSelector:   "article",
}

var goByExample = Source{
	Url:                 "https://gobyexample.com/",
	Domain:              "gobyexample.com",
	ArticleListUrl:      "",
	ArticleListSelector: "li",
	NameSelector:        "a",
	LinkSelector:        "a",
	ArticleMdSelector:   "body",
}

var Sources = []Source{goByExample, flyIo}

