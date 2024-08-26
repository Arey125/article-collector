package article

import (
	"io"
	"net/http"
	"os"
	"path"
)

func (source Source) getHtmlDirectoryPath() string {
	return path.Join(os.Getenv("FILES"), "html", source.Domain)
}

func (source Source) getHtmlPath() string {
	return path.Join(os.Getenv("FILES"), "html", source.Domain, "index.html")
}

func (article Article) getHtmlName() string {
	return article.GetSlug() + ".html"
}

func (article Article) getHtmlPath() string {
	return path.Join(article.Source.getHtmlDirectoryPath(), article.getHtmlName())
}

func getHtml(htmlUrl string, htmlPath string) ([]byte, error) {
	htmlDirectory := path.Dir(htmlPath)
	_, err := os.Stat(htmlPath)

	if err == nil {
		return os.ReadFile(htmlPath)
	}

	if !os.IsNotExist(err) {
		return nil, err
	}

	res, err := http.Get(htmlUrl)
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

func (article Article) getHtml() ([]byte, error) {
	htmlPath := article.getHtmlPath()
	htmlUrl := article.Link
    return getHtml(htmlUrl, htmlPath)
}

func (source Source) getHtml() ([]byte, error) {
	htmlPath := source.getHtmlPath()
	htmlUrl := source.Url + source.ArticleListUrl
    return getHtml(htmlUrl, htmlPath)
}
