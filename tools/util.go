package tools

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func SendResponse(rw http.ResponseWriter, status int, message string, data interface{}) {
	rw.WriteHeader(status)
	response := response{Status: status, Message: message, Data: map[string]interface{}{"data": data}}
	json.NewEncoder(rw).Encode(response)
}

func SendError(rw http.ResponseWriter, status int, err error) {
	SendResponse(rw, status, "error", err.Error())
}

func Decode(rw http.ResponseWriter, r *http.Request, structPointer any) bool {
	if err := json.NewDecoder(r.Body).Decode(structPointer); err != nil {
		SendError(rw, http.StatusBadRequest, err)
		return false
	}
	return true
}

func Validate(rw http.ResponseWriter, structPointer any) bool {
	var validate = validator.New()
	if validationErr := validate.Struct(structPointer); validationErr != nil {
		SendError(rw, http.StatusBadRequest, validationErr)
		return false
	}
	return true
}
