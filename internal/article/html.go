package article

import (
	"io"
	"net/http"
	"os"
	"path"
)
 
func (source Source) GetHtmlDirectoryPath() string {
	return path.Join(os.Getenv("FILES"), "html", source.Domain)
}

func (article Article) getHtmlName() string {
	return article.GetSlug() + ".html"
}

func (article Article) getHtmlPath() string {
	return path.Join(article.Source.GetHtmlDirectoryPath(), article.getHtmlName())
}

func (article Article) GetHtml() ([]byte, error) {
	htmlDirectory := article.Source.GetHtmlDirectoryPath()
	htmlPath := article.getHtmlPath()
	_, err := os.Stat(htmlPath)

	if err == nil {
        return os.ReadFile(htmlPath)
	}

	if !os.IsNotExist(err) {
		return nil, err
	}

	res, err := http.Get(article.Source.Url + article.Link)
	if err != nil {
		return nil, err
	}
    defer res.Body.Close()

    htmlBytes, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }

	err = os.MkdirAll(htmlDirectory, 0o775)
    if err != nil {
        return nil, err
    }

    err = os.WriteFile(htmlPath, htmlBytes, 0o664)
    if err != nil {
        return nil, err
    }

    return htmlBytes, nil
}

