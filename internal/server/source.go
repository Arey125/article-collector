package server

import (
	"context"
	"net/http"

	"github.com/Arey125/article-collector/internal/article"
	"github.com/Arey125/article-collector/internal/server/template"
)

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
		return
	}

	articles, err := server.article.FromSource(source.Id)
	if err != nil {
		serverError(w, err)
	}

	articleLinks := make([]template.TemplateLink, len(articles))
	for i, article := range articles {
        articleLinks[i] = getArticleTemplateLink(article)
	}

	sourcePage := template.SourcePage{
		Title: source.Name,
		Links: articleLinks,
		Nav:   getSourceTemplateNav(*source),
	}

    /*
    templ := NewTemplate("source")
	err = templ.ExecuteTemplate(w, "base", sourcePage)
	if err != nil {
		serverError(w, err)
	}
    */

    page := template.Source(sourcePage)
    page.Render(context.Background(), w)
}
