package main

import (
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
	"io"
)

type genericJSON struct {
	Stuff string `json:"some_stuff"`
}

type textDocument struct {
	Text string `json:"thread_link"`
}


func wrapHandler(
	handler func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {

	h := func(w http.ResponseWriter, r *http.Request) {

		handler(w, r)
	}
	return h
}

func sendErrorMessage(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(status)
	io.WriteString(w, msg)
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	stuff := "Hello goardparser!"
	sendJSONResponse(w, genericJSON{Stuff: stuff})
}

func parseDataHandler(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read the request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var td textDocument

	if err := json.Unmarshal(body, &td); err != nil {
		sendErrorMessage(w, "Could not decode the request body as JSON", http.StatusBadRequest)
		return
	}
	if td.Text == ""  {
		sendErrorMessage(
			w,
			"Required parameter {thread_link} is missing",
			http.StatusBadRequest)
		return
	}
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