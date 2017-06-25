package main

import (
	"net/http"
	"os"
	"github.com/rs/cors"
)


func main() {
	r := MakeRouter()
	http.Handle("/", r)

	//handler := cors.New(cors.Options{
	//	AllowedOrigins: []string{"https://goardparcerface.herokuapp.com"},
	//}).Handler(r)
	handler := cors.Default().Handler(r)

	http.ListenAndServe(
		":"+os.Getenv("PORT"),
		handler)
}