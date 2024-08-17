package server

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/Arey125/article-collector/internal/article"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

var mdRenderer = goldmark.New(goldmark.WithExtensions(highlighting.Highlighting))

func (server *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", server.Home)
	mux.HandleFunc("GET /blog/{path...}", server.Blog)

	return mux
}

func (server *Server) Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Home page\n")
	for _, source := range article.Sources {
		articles, _ := source.GetArticleList()
		for _, article := range articles {
			fmt.Fprintf(w, "%s/%s\n", article.Source.Domain, path.Base(article.Link))
		}
	}
}

func (server *Server) Blog(w http.ResponseWriter, req *http.Request) {
	path := path.Join(os.Getenv("FILES"), req.PathValue("path")+".md")

	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprint(w, "No such file")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	mdRenderer.Convert(file, w)
}
