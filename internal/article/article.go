package article

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

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

func (article Article) GetSlug() string {
    return path.Base(article.Link)
}

func (article Article) getFileName() string {
	return article.GetSlug() + ".md"
}

func (article Article) GetFilePath() string {
	return path.Join(article.Source.GetDirectoryPath(), article.getFileName())
}

func (article Article) SaveToFileIfDoesNotExist() (saved bool, error error) {
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

	md, err := article.GetMd()
	if err != nil {
		return false, err
	}

	outputFile.Write([]byte(md))
	return true, nil
}

func SaveAllArticlesFromSource(source Source) error {
	articles, err := source.GetArticleList()
	if err != nil {
        return err
	}

	for _, article := range articles {
		saved, err := article.SaveToFileIfDoesNotExist()
        if err != nil {
            panic(err)
        }
		if saved {
            fmt.Printf("\"%s\" from %s is saved\n", article.Name, article.Source.Domain)
			time.Sleep(400 * time.Millisecond)
		}
	}

    return nil
}
