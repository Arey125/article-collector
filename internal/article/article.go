package article

import (
	"bytes"
	"os"
	"path"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Name   string
	Link   string
	Slug   string
	Source *Source

	Status string
    Sort   int
}

func NewArticle(name string, link string, source *Source) Article {
    return Article{
        Name:   name,
        Link:   link,
        Source: source,
        Slug:   path.Base(link),

        Status: "unread",
    }
}

func (article Article) GetSlug() string {
	return path.Base(article.Link)
}

func (article Article) getFileName() string {
	return article.GetSlug() + ".md"
}

func (article Article) GetFilePath() string {
	return path.Join(article.Source.GetDirectoryPath(), article.getFileName())
}

func (article Article) GetMd() (string, error) {
	html, err := article.getHtml()
	if err != nil {
		return "", err
	}
	htmlBuffer := bytes.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(htmlBuffer)
	if err != nil {
		return "", err
	}

	articleElement := doc.Find(article.Source.ArticleMdSelector)
	converter := md.NewConverter(article.Source.Domain, true, nil)

	return converter.Convert(articleElement), nil
}

func (article Article) SaveToFileIfDoesNotExist() (isFromNetwork bool, error error) {
	directory := article.Source.GetDirectoryPath()
	path := article.GetFilePath()

	err := os.MkdirAll(directory, 0o775)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(path)
	if !os.IsNotExist(err) {
		return false, err
	}
	if err == nil {
		return false, nil
	}

	outputFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	htmlPath := article.getHtmlPath()
	_, htmlFileErr := os.Stat(htmlPath)

	md, err := article.GetMd()
	if err != nil {
		return false, err
	}

	outputFile.Write([]byte(md))
	return htmlFileErr != nil, nil
}
