package main

import "github.com/gorilla/mux"

func MakeRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/text", wrapHandler(indexHandler)).Methods("GET")
	r.HandleFunc("/text/{hash}", wrapHandler(indexHandler)).Methods("GET")

	return r
}
