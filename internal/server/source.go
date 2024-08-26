package server

import (
	"html/template"
	"net/http"

	"github.com/Arey125/article-collector/internal/article"
)

type SourcePage struct {
	Title string
    Nav   []Link
	Links []Link
}

func (server *Server) Source(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	sourceDomain := req.PathValue("source")

	var source *article.Source = nil
	for _, cur := range article.Sources {
		if cur.Domain == sourceDomain {
			source = cur
			break
		}
	}

	if source == nil {
        notFound(w)
        return;
	}

	articles, err := server.article.FromSource(source.Id)
    if err != nil {
        serverError(w, err);
    }

	articleLinks := make([]Link, len(articles))
    for i, article := range articles {
        articleLinks[i] = getArticleLink(article)
    }

	sourcePage := SourcePage{
		Title: source.Name,
        Links: articleLinks,
        Nav:   getSourceNav(*source),
	}

	templ := template.Must(template.ParseFiles("ui/base.html", "ui/pages/source.html", "ui/partials/nav.html"))
    err = templ.ExecuteTemplate(w, "base", sourcePage)
    if err != nil {
        serverError(w, err)
    }
}
