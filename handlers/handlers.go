package handlers

import (
	"goardparser/structs"
	"goardparser/errors"
	"goardparser/validators"
	"goardparser/utils"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	stuff := "Hello goardparser!"
	utils.JSONResponse(w, structs.GenericJSON{Stuff: stuff})
}

func ParseDataHandler(w http.ResponseWriter, r *http.Request){
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
	if validators.IsValidRequestParams(w, td) {

		ch := make(chan *structs.Result)
		go parseThread(td.Data, ch)
		data :=  <-ch

		responseJson := &structs.ResponseJSON{}

		for _, item := range data.Threads[0].Posts {

			for _, file := range item.Files{

				if strings.Contains(file.Name, ".webm"){
					responseJson.Files = append(responseJson.Files, file)
				}
			}
		}
		utils.JSONResponse(w, responseJson)
	}

}

func parseThread(url string, ch chan *structs.Result) {
	log.Printf("Making request to: %v", url)
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}


	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	result := &structs.Result{}
	json.Unmarshal(body, result)

	ch <- result

}