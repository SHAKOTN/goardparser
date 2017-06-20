package validators

import (
	"net/http"
	"goardparser/structs"
	"goardparser/utils"
	"strings"
)

func IsValidRequestParams(w http.ResponseWriter, st structs.RequestDataJSON) bool{
	if st.Data == ""  {
		utils.JSONResponse(
			w,
			structs.ErrorMsg{Msg: "Required parameter {thread_link} is missing"},
			http.StatusBadRequest)
		return false
	}

	if !strings.Contains(st.Data, "https://2ch.hk/"){
		utils.JSONResponse(
			w,
			structs.ErrorMsg{Msg: "Required parameter {thread_link} is missing"},
			http.StatusNotAcceptable)
		return false
	}
	return true
}