package server

import (
	"html/template"
	"net/http"
	"path"

	"github.com/Arey125/article-collector/internal/article"
)

func (server *Server) Home(w http.ResponseWriter, req *http.Request) {
	blogLinks := make([]Link, len(article.Sources))

    for i, blog := range article.Sources {
        blogLinks[i] = Link{
            Title: blog.Domain,
            Link: path.Join("/blog", blog.Domain),
        }
    }

	templ := template.Must(template.ParseFiles("internal/server/home.html"))
	templ.Execute(w, blogLinks)
}
