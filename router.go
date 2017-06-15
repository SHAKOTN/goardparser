package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MakeRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", wrapHandler(indexHandler)).Methods("GET")
	r.HandleFunc("/parse_data", wrapHandler(parseDataHandler)).Methods("POST")

	return r
}

func wrapHandler(
	handler func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {

	h := func(w http.ResponseWriter, r *http.Request) {

		handler(w, r)
	}
	return h
}