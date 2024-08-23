package server

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func serverError(w http.ResponseWriter, err error) {
    trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
    log.Output(2, trace)

    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func clientError(w http.ResponseWriter, status int) {
    http.Error(w, http.StatusText(status), status)
}

func notFound(w http.ResponseWriter) {
    clientError(w, http.StatusNotFound)
}
