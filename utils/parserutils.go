package utils

import (
	"net/http"
	"encoding/json"
	"log"
	"goardparser/structs"
	"io/ioutil"
	"errors"
)

func JSONResponse(w http.ResponseWriter, data interface{}) {
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

func ParseThread(url string, ch chan *structs.Result) {
	log.Printf("Making request to: %v", url)
	res, err := http.Get(url)

	if err != nil {
		log.Print(err)
	}

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		errorStr := string(b)
		log.Printf("Cannot obtain thread")

		result := &structs.Result{}
		result.Error = errors.New(errorStr)

		ch <- result

	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Print(err)
		}

		result := &structs.Result{}
		json.Unmarshal(body, result)

		ch <- result
	}


}