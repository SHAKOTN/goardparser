package validators

import (
	"net/http"
	"goardparser/structs"
	"goardparser/utils"
	"strings"
)

type Validator struct {
	Origin string
}

func (v *Validator) IsValidRequestParams(w http.ResponseWriter, st structs.RequestDataJSON) bool{
	if st.Data == ""  {
		utils.JSONResponse(
			w,
			structs.ErrorMsg{Msg: "Required parameter {thread_link} is missing"},
			http.StatusBadRequest)
		return false
	}
	if strings.Contains(st.Data, "https://2ch.hk/") {
		v.Origin = "https://2ch.hk"
	// TODO: implement and test it
	} else if strings.Contains(st.Data, "http://www.4chan.org/") {
		v.Origin = "http://www.4chan.org"
	} else {
		utils.JSONResponse(
			w,
			structs.ErrorMsg{Msg: "This is not an imageboard link"},
			http.StatusNotAcceptable)
		return false
	}

	return true
}