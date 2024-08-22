package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/Arey125/article-collector/internal/article"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

var mdRenderer = goldmark.New(
	goldmark.WithExtensions(
		highlighting.NewHighlighting(
			highlighting.WithStyle("catppuccin"),
			highlighting.WithGuessLanguage(true),
		),
	),
)

type ArticlePage struct {
	Title   string
	Source  Link
	Content template.HTML
}

func (server *Server) Article(w http.ResponseWriter, req *http.Request) {
	sourceDomain := req.PathValue("source")
	articleSlug := req.PathValue("article")

	var source *article.Source = nil
	for _, cur := range article.Sources {
		if cur.Domain == sourceDomain {
			source = &cur
			break
		}
	}

	var currentArticle *article.Article = nil
    articleList, err := source.GetArticleList()
    if err != nil {
		fmt.Fprint(w, "Cannot get source article list")
        return;
    }
	for _, cur := range articleList {
		if cur.GetSlug() == articleSlug {
			currentArticle = &cur
			break
		}
	}

	filePath := currentArticle.GetFilePath()

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprint(w, "No such file")
		return
	}

	contentBuffer := bytes.NewBuffer(make([]byte, 0))
	mdRenderer.Convert(file, contentBuffer)

	articlePage := ArticlePage{
		Title: articleSlug,
		Source: getSourceLink(*source),
		Content: template.HTML(contentBuffer.String()),
	}

	templ := template.Must(template.ParseFiles("internal/server/article.html"))
	templ.Execute(w, articlePage)
}
