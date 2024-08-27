package server

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"

	"github.com/Arey125/article-collector/internal/article"
	. "github.com/Arey125/article-collector/internal/server/template"
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
    Status  string
	Nav     []Link
	Content template.HTML
}

func (server *Server) Article(w http.ResponseWriter, req *http.Request) {
	sourceDomain := req.PathValue("source")
	articleSlug := req.PathValue("article")

	var source *article.Source = nil
	for _, cur := range article.Sources {
		if cur.Domain == sourceDomain {
			source = cur
			break
		}
	}
    if (source == nil) {
        notFound(w);
        return;
    }

	var currentArticle *article.Article = nil
	articleList, err := source.GetArticleListFromHtml()
	if err != nil {
        serverError(w, err);
		return
	}
	for _, cur := range articleList {
		if cur.GetSlug() == articleSlug {
			currentArticle = &cur
			break
		}
	}
    if (currentArticle == nil) {
        notFound(w);
        return;
    }

	filePath := currentArticle.GetFilePath()

	file, err := os.ReadFile(filePath)
	if err != nil {
		notFound(w)
		return
	}

	contentBuffer := bytes.NewBuffer(make([]byte, 0))
	mdRenderer.Convert(file, contentBuffer)

	articlePage := ArticlePage{
		Title:   currentArticle.Name,
		Nav: getArticleNav(*currentArticle),
        Content: template.HTML(contentBuffer.String()),
        Status: currentArticle.Status,
	}

	templ := NewTemplate("article")
    err = templ.ExecuteTemplate(w, "base", articlePage)
    if err != nil {
        serverError(w, err)
    }
}
