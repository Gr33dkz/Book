package http

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type AcceptedResp struct {
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"Message,omitempty"`
}

type ErrorMessage struct {
	StatusCode   int    `json:"StatusCode"`
	ErrorMessage string `json:"ErrorMessage"`
}

func Accepted(
	writer http.ResponseWriter,
	message string) {

	m := http.StatusText(http.StatusOK)
	if message != "" {
		m = message
	}

	msg := AcceptedResp{
		StatusCode: http.StatusOK,
		Message:    m,
	}

	rsp, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("got marshaling error", err) // TODO ADD LOGGER
	}
	setDefaultHeaders(writer)
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(rsp)
	if err != nil {
		fmt.Println("got write error") // TODO ADD LOGGER
	}
}

func Response(
	w http.ResponseWriter,
	entity interface{}) {
	rsp, err := json.Marshal(entity)
	if err != nil {
		MarshallError(w)
	}
	setDefaultHeaders(w)
	w.WriteHeader(http.StatusOK)
	if entity != nil {
		_, err = w.Write(rsp)
		if err != nil {
			fmt.Println("got write error") // TODO ADD LOGGER
		}
	}

}

func ResponseWithError(
	writer http.ResponseWriter,
	statusCode int,
	errorMessage string) {

	msg := ErrorMessage{
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}
	rsp, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("got marshaling error", err) // TODO ADD LOGGER
	}
	setDefaultHeaders(writer)
	writer.WriteHeader(statusCode)
	_, err = writer.Write(rsp)
	if err != nil {
		fmt.Println("got write error") // TODO ADD LOGGER
	}
}

func setDefaultHeaders(w http.ResponseWriter) {
	for k, v := range DefaultHeaders {
		w.Header().Set(k, v)
	}
}

func UnmarshallError(w http.ResponseWriter) {
	ResponseWithError(w, http.StatusBadRequest, "unmarshall error")
}

func MarshallError(w http.ResponseWriter) {
	ResponseWithError(w, http.StatusBadRequest, "marshall error")
}

func AlreadyExist(w http.ResponseWriter) {
	ResponseWithError(w, http.StatusConflict, "entity already exists")
}

func GenerateId() string {
	return uuid.NewString()
}
