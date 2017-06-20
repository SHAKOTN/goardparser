package utils

import (
	"net/http"
	"encoding/json"
	"log"
	"goardparser/structs"
	"io/ioutil"
	"errors"
)

func JSONResponse(writer http.ResponseWriter, data interface{}, status int) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(status)
	_, err = writer.Write(body)
	if err != nil {
		log.Printf("Failed to write the response body: %v", err)
		return
	}
}


func ParseThread(url string, ch chan *structs.Board) {
	log.Printf("Making request to: %v", url)
	res, err := http.Get(url)

	if err != nil {
		log.Print(err)
	}

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		errorStr := string(b)
		log.Printf("Cannot obtain thread")

		result := &structs.Board{}
		result.Error = errors.New(errorStr)

		ch <- result

	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Print(err)
		}

		result := &structs.Board{}
		json.Unmarshal(body, result)

		ch <- result
	}
}