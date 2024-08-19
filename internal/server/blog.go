package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/Arey125/article-collector/internal/article"
)

type Link struct {
	Title string
	Link  string
}

type BlogPage struct {
	Title string
	Links []Link
}

func (server *Server) Blog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	blogDomain := req.PathValue("blog")

	var blog *article.Source = nil
	for _, cur := range article.Sources {
		if cur.Domain == blogDomain {
			blog = &cur
			break
		}
	}

	if blog == nil {
		w.Write([]byte("No such blog"))
	}

	articles, err := blog.GetArticleList()
    if err != nil {
        panic(err)
    }

	articleLinks := make([]Link, len(articles))
    for i, article := range articles {
        articleLinks[i] = Link{
            Title: article.Name,
            Link: fmt.Sprintf("/blog/%s/%s", article.Source.Domain, path.Base(article.Link)),
        }
    }

	blogPage := BlogPage{
		Title: blogDomain,
        Links: articleLinks,
	}

	templ := template.Must(template.ParseFiles("internal/server/blog.html"))
	templ.Execute(w, blogPage)
}
