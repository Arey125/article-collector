package article

import (
	"bytes"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Source struct {
	Id     string
	Url    string
	Domain string
	Name   string

	ArticleListUrl      string
	ArticleListSelector string
	NameSelector        string
	LinkSelector        string
	ArticleMdSelector   string
}

func (source Source) GetArticleListFromHtml() ([]Article, error) {
	html, err := source.getHtml()
	if err != nil {
		return nil, err
	}
    htmlBuffer := bytes.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(htmlBuffer)
	if err != nil {
		return nil, err
	}

	articleSelect := doc.Find(source.ArticleListSelector)
	articles := make([]Article, len(articleSelect.Nodes), len(articleSelect.Nodes))
	for i := range articleSelect.Nodes {
		articleNode := articleSelect.Eq(i)
		articles[i] = Article{
			Name:   strings.TrimSpace(articleNode.Find(source.NameSelector).Text()),
			Link:   source.Url + articleNode.Find(source.LinkSelector).AttrOr("href", ""),
			Source: &source,
		}
	}

	return articles, nil
}

func (source Source) GetDirectoryPath() string {
	return path.Join(os.Getenv("FILES"), "md" ,source.Domain)
}

var flyIo = Source{
	Id:     "fly.io",
	Url:    "https://fly.io",
	Domain: "fly.io",
	Name:   "Fly.io",

	ArticleListUrl:      "/blog",
	ArticleListSelector: "article",
	NameSelector:        "h1",
	LinkSelector:        "a",
	ArticleMdSelector:   "article",
}

var goByExample = Source{
	Id:     "gobyexample.com",
	Url:    "https://gobyexample.com/",
	Domain: "gobyexample.com",
	Name:   "Go by Example",

	ArticleListUrl:      "",
	ArticleListSelector: "li",
	NameSelector:        "a",
	LinkSelector:        "a",
	ArticleMdSelector:   "table, p",
}

var Sources = []*Source{&goByExample, &flyIo}
var SourceMap = (func () map[string]*Source {
    sourcesMap := map[string]*Source{}
    for _, source := range Sources {
        sourcesMap[source.Id] = source
    }
    return sourcesMap
})()
