package server

import (
	"net/http"

	"github.com/Arey125/article-collector/internal/article"
    . "github.com/Arey125/article-collector/internal/server/template"
)

func (server *Server) Home(w http.ResponseWriter, req *http.Request) {
	sourceLinks := make([]Link, len(article.Sources))

    for i, source := range article.Sources {
        sourceLinks[i] = getSourceLink(*source)
    }

	templ := NewTemplate("home")
	templ.ExecuteTemplate(w, "base", sourceLinks)
}
