package article

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Source struct {
	Url    string
	Domain string
	Name   string
	Id     string

	ArticleListUrl      string
	ArticleListSelector string
	NameSelector        string
	LinkSelector        string
	ArticleMdSelector   string
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

func (source Source) GetDirectoryPath() string {
	return path.Join(os.Getenv("FILES"), source.Domain)
}

var flyIo = Source{
	Url:    "https://fly.io",
	Domain: "fly.io",
	Name:   "Fly.io",
	Id:     "fly.io",

	ArticleListUrl:      "/blog",
	ArticleListSelector: "article",
	NameSelector:        "h1",
	LinkSelector:        "a",
	ArticleMdSelector:   "article",
}

var goByExample = Source{
	Url:    "https://gobyexample.com/",
	Domain: "gobyexample.com",
	Name:   "Go by Example",
	Id:     "gobyexample.com",

	ArticleListUrl:      "",
	ArticleListSelector: "li",
	NameSelector:        "a",
	LinkSelector:        "a",
	ArticleMdSelector:   "table, p",
}

var Sources = []Source{goByExample, flyIo}
