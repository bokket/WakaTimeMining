package router

import (
	"net/http"

	"wakever/handler"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {

	// serve static file request
	fs := http.FileServer(http.Dir("assets/"))
	serveFileHandler := http.StripPrefix("/static/", fs)

	r.PathPrefix("/static/").Handler(serveFileHandler)

    viewRouter := r.PathPrefix("/view").Subrouter()
    viewRouter.HandleFunc("/index", handler.ShowIndexView)
}