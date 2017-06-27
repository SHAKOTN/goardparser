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
	"os"
	"io"
	"path/filepath"
)

func IndexHandler(writer http.ResponseWriter, r *http.Request)  {
	stuff := "Hello goardparser!"
	utils.JSONResponse(writer, structs.GenericJSON{Stuff: stuff}, http.StatusOK)
}

func ParseDataHandler(writer http.ResponseWriter, request *http.Request){
	body, err := ioutil.ReadAll(request.Body)
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

		data := utils.ParseThread(requestData.Data)

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
					// TODO: This is very ugly. Refactor
					file.NormalizeSrcPath(validator.Origin)
					responseJson.Files = append(responseJson.Files, file)
				}
			}
		}
		utils.JSONResponse(writer, responseJson, http.StatusOK)
	}
}


func DownloadDataHandler(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Failed to read the request body: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var requestData structs.InnerPost

	if err := json.Unmarshal(body, &requestData); err != nil {
		utils.JSONResponse(writer,
			structs.ErrorMsg{Msg: "Could not decode the request body as JSON"},
			http.StatusBadRequest)
		return
	}
	if len(requestData.Files) > 50 {
		utils.JSONResponse(writer,
			structs.ErrorMsg{Msg: "Too many files to download"},
			http.StatusBadRequest)
		return
	}
	path := "tmp/"
	err = os.MkdirAll(path, os.FileMode(0777))

	if err != nil {
		log.Println("Error creating directory")
		log.Println(err)
		return
	}
	var tasks []*utils.Task

	for _, file := range requestData.Files {
		var fileName = filepath.Join(path, file.Name)
		var filePath = file.Path

		tasks = append(tasks, utils.NewTask(func() error {
			log.Printf("Downloading %v", fileName)
			out, err := os.Create(fileName)
			if err != nil{
				log.Printf("Can't create file")
			}
			defer out.Close()
			resp, err := http.Get(filePath)
			if err != nil{
				log.Printf("Can't download file")
			}
			defer resp.Body.Close()

			_, err = io.Copy(out, resp.Body)

			if err != nil{
				log.Print(err)
			}
			return err
		}))
	}

	runTasks(tasks)
}

func runTasks(tasks []*utils.Task) bool {
	p := utils.NewPool(tasks, 50)
	p.Run()
	var numErrors int
	for _, task := range p.Tasks {
		if task.Err != nil {
			log.Print(task.Err)
			numErrors++
		}
		if numErrors >= 10 {
			log.Print("Too many errors.")
			break
		}
	}

	return !p.HasErrors()
}