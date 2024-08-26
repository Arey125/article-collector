package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Arey125/article-collector/internal/models"
)

type Server struct {
    port int
    article models.ArticleModel
}

func NewServer(db *sql.DB) *http.Server {
    port, _ := strconv.Atoi(os.Getenv("PORT"))
    newServer := &Server {
        port: port,
        article: models.ArticleModel{DB: db},
    }

    server := &http.Server {
        Addr: fmt.Sprintf(":%d", newServer.port),
        Handler: newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
    }

    return server
}
