package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"goardparser/handlers"
)

func MakeRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", wrapHandler(handlers.IndexHandler)).Methods("GET")
	router.HandleFunc("/parse_data", wrapHandler(handlers.ParseDataHandler)).Methods("POST")
	//router.HandleFunc("/download", wrapHandler(handlers.DownloadDataHandler)).Methods("POST")

	return router
}

func wrapHandler(
	handler func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {

	inner_handler := func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
	return inner_handler
}