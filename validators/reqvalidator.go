package validators

import (
	"net/http"
	"goardparser/structs"
	"goardparser/errors"
	"strings"
)

func ValidateParseHandlerParams(w http.ResponseWriter, st structs.RequestDataJSON){
	if st.Data == ""  {
		errors.SendErrorMessage(
			w,
			"Required parameter {thread_link} is missing",
			http.StatusBadRequest)
		return
	}

	if !strings.Contains(st.Data, "http://2ch.hk/"){
		errors.SendErrorMessage(
			w,
			"The link is invalid. You should use http(s)://link-to-image-board",
			http.StatusNotAcceptable)
		return
	}
}