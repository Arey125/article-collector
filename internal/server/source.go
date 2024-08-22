package server

import (
	"html/template"
	"net/http"

	"github.com/Arey125/article-collector/internal/article"
)

type SourcePage struct {
	Title string
	Links []Link
}

func (server *Server) Source(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	sourceDomain := req.PathValue("source")

	var source *article.Source = nil
	for _, cur := range article.Sources {
		if cur.Domain == sourceDomain {
			source = &cur
			break
		}
	}

	if source == nil {
		w.Write([]byte("No such source"))
	}

	articles, err := source.GetArticleList()
    if err != nil {
        panic(err)
    }

	articleLinks := make([]Link, len(articles))
    for i, article := range articles {
        articleLinks[i] = getArticleLink(article)
    }

	sourcePage := SourcePage{
		Title: source.Name,
        Links: articleLinks,
	}

	templ := template.Must(template.ParseFiles("ui/source.html"))
	templ.Execute(w, sourcePage)
}
