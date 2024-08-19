package server

import (
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

func (server *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", server.Home)
	mux.HandleFunc("GET /blog/{path...}", server.Blog)

	return mux
}

func (server *Server) Home(w http.ResponseWriter, req *http.Request) {
	allArticles := make([]article.Article, 0)
	for _, source := range article.Sources {
		articles, _ := source.GetArticleList()
		allArticles = append(allArticles, articles...)
	}
	type ArticleLink struct {
		Title string
		Link  string
	}

	articleLinks := make([]ArticleLink, len(allArticles))
	for i, article := range allArticles {
		articleLinks[i].Title = fmt.Sprintf("%s from %s", article.Name, article.Source.Domain)
		articleLinks[i].Link = fmt.Sprintf("/blog/%s/%s", article.Source.Domain, path.Base(article.Link))
	}

	templ := template.Must(template.ParseFiles("internal/server/home.html"))
	templ.Execute(w, articleLinks)
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
