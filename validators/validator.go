package validators

import (
	"net/http"
	"goardparser/structs"
	"goardparser/errors"
	"strings"
)

func IsValidRequestParams(w http.ResponseWriter, st structs.RequestDataJSON) bool{
	if st.Data == ""  {
		errors.SendErrorMessage(
			w,
			"Required parameter {thread_link} is missing",
			http.StatusBadRequest)
		return false
	}

	if !strings.Contains(st.Data, "https://2ch.hk/"){
		errors.SendErrorMessage(
			w,
			"The link is invalid. You should use http(s)://link-to-image-board",
			http.StatusNotAcceptable)
		return false
	}
	return true
}