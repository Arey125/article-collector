package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"

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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	sourceDomain := req.PathValue("source")
	articleSlug := req.PathValue("article")

	var source *article.Source = nil
	for _, cur := range article.Sources {
		if cur.Domain == sourceDomain {
			source = &cur
			break
		}
	}

	filePath := path.Join(os.Getenv("FILES"), sourceDomain, articleSlug+".md")

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprint(w, "No such file")
		return
	}

	contentBuffer := bytes.NewBuffer(make([]byte, 0))
	mdRenderer.Convert(file, contentBuffer)

	articlePage := ArticlePage{
		Title: articleSlug,
		Source: Link{
			Title: source.Name,
			Link:  path.Join("/source", sourceDomain),
		},
		Content: template.HTML(contentBuffer.String()),
	}

	templ := template.Must(template.ParseFiles("internal/server/article.html"))
	templ.Execute(w, articlePage)
}
