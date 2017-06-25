package handlers

import (
	"goardparser/structs"
	"goardparser/validators"
	"goardparser/utils"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
	"strings"
)

func IndexHandler(writer http.ResponseWriter, r *http.Request)  {
	stuff := "Hello goardparser!"
	utils.JSONResponse(writer, structs.GenericJSON{Stuff: stuff}, http.StatusOK)
}

func ParseDataHandler(writer http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read the request body: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var requestData structs.RequestDataJSON

	if err := json.Unmarshal(body, &requestData); err != nil {
		utils.JSONResponse(writer,
			structs.ErrorMsg{Msg: "Could not decode the request body as JSON"},
			http.StatusBadRequest)
		return
	}
	validator := validators.Validator{}
	if validator.IsValidRequestParams(writer, requestData) {

		channel := make(chan *structs.Board)
		go utils.ParseThread(requestData.Data, channel)

		data :=  <-channel

		if data.Error != nil || len(data.Threads) == 0{
			utils.JSONResponse(writer,
				structs.ErrorMsg{Msg: "Thread does not exist or it is empty"},
				http.StatusBadRequest)
			return
		}
		responseJson := &structs.ResponseJSON{}

		for _, post := range data.Threads[0].Posts {

			for _, file := range post.Files{

				if strings.Contains(file.Name, ".webm"){
					file.NormalizeSrcPath(validator.Origin)
					responseJson.Files = append(responseJson.Files, file)
				}
			}
		}
		utils.JSONResponse(writer, responseJson, http.StatusOK)
	}
}
