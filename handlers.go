package main

import (
	"goardparser/structs"
	"goardparser/errors"
	"goardparser/validators"
	"goardparser/utils"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
)


func indexHandler(w http.ResponseWriter, r *http.Request)  {
	stuff := "Hello goardparser!"
	utils.JSONResponse(w, structs.GenericJSON{Stuff: stuff})
}

func parseDataHandler(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read the request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var td structs.RequestDataJSON

	if err := json.Unmarshal(body, &td); err != nil {
		errors.SendErrorMessage(w, "Could not decode the request body as JSON", http.StatusBadRequest)
		return
	}
	validators.ValidateParseHandlerParams(w, td)
}