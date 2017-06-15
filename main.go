package main

import (
	"net/http"
	"os"
)


func main() {
	r := MakeRouter()
	http.Handle("/", r)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}