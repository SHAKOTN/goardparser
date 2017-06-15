package main

import "github.com/gorilla/mux"

func MakeRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", wrapHandler(indexHandler)).Methods("GET")

	return r
}
