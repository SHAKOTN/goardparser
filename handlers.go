package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type genericJSON struct {
	Stuff string `json:"some_stuff"`
}

func wrapHandler(
	handler func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {

	h := func(w http.ResponseWriter, r *http.Request) {

		handler(w, r)
	}
	return h
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	stuff := "Hello goardparser!"
	sendJSONResponse(w, genericJSON{Stuff: stuff})
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write the response body: %v", err)
		return
	}
}