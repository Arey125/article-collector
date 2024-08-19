package server

import (
	"fmt"
	"net/http"
	"os"
	"path"

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

func (server *Server) Article(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	blog := req.PathValue("blog")
	articleSlug := req.PathValue("article")

	path := path.Join(os.Getenv("FILES"), blog, articleSlug+".md")

	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprint(w, "No such file")
		return
	}

	mdRenderer.Convert(file, w)
}
