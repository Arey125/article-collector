package server

import (
	"html/template"
	"net/http"
)

func (server *Server) Status(w http.ResponseWriter, req *http.Request) {
	sourceDomain := req.PathValue("source")
	articleSlug := req.PathValue("article")
	action := req.PathValue("action")

	if action != "read" && action != "unread" {
		clientError(w, http.StatusBadRequest)
		return
	}

	article, err := server.article.Get(sourceDomain, articleSlug)
	if err != nil {
		serverError(w, err)
	}

	tmpl := template.Must(template.ParseFiles("./ui/partials/status.html"))
	err = tmpl.ExecuteTemplate(w, "status", article.Status)
	if err != nil {
        serverError(w, err)
	}
}
