package server

import (
	"net/http"
)

func (server *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", server.Home)
	mux.HandleFunc("GET /blog/{blog}", server.Blog)
	mux.HandleFunc("GET /blog/{blog}/{article}", server.Article)

	return mux
}
