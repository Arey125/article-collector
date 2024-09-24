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
    Article article.Article
	Nav     []Link
	Content template.HTML
}

func (server *Server) Article(w http.ResponseWriter, req *http.Request) {
	sourceDomain := req.PathValue("source")
	articleSlug := req.PathValue("article")

    currentArticle, err := server.article.Get(sourceDomain, articleSlug)
    if err != nil {
        notFound(w)
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
        Article: *currentArticle,
	}

	templ := NewTemplate("article")
    err = templ.ExecuteTemplate(w, "base", articlePage)
    if err != nil {
        serverError(w, err)
    }
}
