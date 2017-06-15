package main

import (
	"net/http"
)


func main() {
	r := MakeRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}