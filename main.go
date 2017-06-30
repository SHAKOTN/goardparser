package main

import (
	"net/http"
	"os"
	"github.com/rs/cors"
	"runtime"
)


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := MakeRouter()
	http.Handle("/", r)

	var handler http.Handler

	if os.Getenv("ENV") == "development" {
		handler = cors.Default().Handler(r)
	} else {
		handler = cors.New(cors.Options{
			AllowedOrigins: []string{"https://goardparcerface.herokuapp.com"},
		}).Handler(r)
	}

	http.ListenAndServe(
		":"+os.Getenv("PORT"),
		handler)
}