package server

import (
	"context"
	"net/http"

	"github.com/Arey125/article-collector/internal/article"
	"github.com/Arey125/article-collector/internal/server/template"
)

func (server *Server) Home(w http.ResponseWriter, req *http.Request) {
	sourceLinks := make([]template.TemplateLink, len(article.Sources))

    for i, source := range article.Sources {
        link := getSourceLink(*source)
        sourceLinks[i] = template.TemplateLink{
            Title: link.Title,
            Link: link.Link,
        }
    }

    page := template.Home(sourceLinks)
    page.Render(context.Background(), w)
}
