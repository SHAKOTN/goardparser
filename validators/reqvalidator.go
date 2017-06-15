package validators

import (
	"goardparser/structs"
	"goardparser/errors"
	"net/http"
)

func ValidateParseHandlerParams(w http.ResponseWriter, st structs.RequestDataJSON){
	if st.Data == ""  {
		errors.SendErrorMessage(
			w,
			"Required parameter {thread_link} is missing",
			http.StatusBadRequest)
		return
	}
}