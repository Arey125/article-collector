package server

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"text/template"

	"github.com/Arey125/article-collector/internal/article"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

const homeTemplateString = `
<html>
    <head>
        <title>Home</title>
    </head>
    <body>
        <h1>Home</h1>
        <ul>
            {{ range . }}
            <li><a href="/blog/{{ . }}">{{ . }}</a></li>
            {{ end }}
        </ul>
    </body>
</html>
`

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

	articleStrings := make([]string, len(allArticles))
	for i, article := range allArticles {
		articleStrings[i] = fmt.Sprintf("%s/%s", article.Source.Domain, path.Base(article.Link))
	}

	templ, err := template.New("home").Parse(homeTemplateString)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templ.Execute(w, articleStrings)
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
