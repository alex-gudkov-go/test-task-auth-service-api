package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func Write(rw http.ResponseWriter, data any, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(data)
}

func WriteOK(rw http.ResponseWriter, data any) {
	Write(rw, data, http.StatusOK)
}

func WriteInternalServerError(rw http.ResponseWriter, message string) {
	Write(rw, Error{Message: message}, http.StatusInternalServerError)
}

func WriteBadRequest(rw http.ResponseWriter, message string) {
	Write(rw, Error{Message: message}, http.StatusBadRequest)
}

func WriteUnauthorized(rw http.ResponseWriter, message string) {
	Write(rw, Error{Message: message}, http.StatusUnauthorized)
}
