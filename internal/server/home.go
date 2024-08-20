package server

import (
	"html/template"
	"net/http"

	"github.com/Arey125/article-collector/internal/article"
)

func (server *Server) Home(w http.ResponseWriter, req *http.Request) {
	sourceLinks := make([]Link, len(article.Sources))

    for i, source := range article.Sources {
        sourceLinks[i] = getSourceLink(source)
    }

	templ := template.Must(template.ParseFiles("internal/server/home.html"))
	templ.Execute(w, sourceLinks)
}
