package server

import (
	"net/http"
)

func (server *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", server.Home)
	mux.HandleFunc("GET /source/{source}", server.Source)
	mux.HandleFunc("GET /source/{source}/{article}", server.Article)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	return mux
}
