package article

import (
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Name   string
	Link   string
	Source *Source
}

func (article Article) GetMd() (string, error) {
	res, err := http.Get(article.Source.Url + article.Link)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	articleElement := doc.Find(article.Source.ArticleMdSelector)
	converter := md.NewConverter(article.Source.Domain, true, nil)

	return converter.Convert(articleElement), nil
}
