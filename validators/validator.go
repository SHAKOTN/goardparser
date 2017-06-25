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
	if strings.HasPrefix(st.Data, "https://2ch.hk/") {
		v.Origin = "https://2ch.hk"
	} else if strings.HasPrefix(st.Data, "http://www.4chan.org/") {
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