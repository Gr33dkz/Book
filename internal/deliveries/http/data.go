package http

import (
	"encoding/json"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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
		log.Debug("got marshaling error")
	}
	setDefaultHeaders(writer)
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(rsp)
	if err != nil {
		log.Debug("got write error")
	}
}

func Response(
	w http.ResponseWriter,
	entity interface{}) {
	rsp, err := json.Marshal(entity)
	if err != nil {
		log.Debug("got marshal error")
	}
	setDefaultHeaders(w)
	w.WriteHeader(http.StatusOK)
	if entity != nil {
		_, err = w.Write(rsp)
		if err != nil {
			log.Debug("got write error")
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
		log.Debug("got marshaling error")
	}
	setDefaultHeaders(writer)
	writer.WriteHeader(statusCode)
	_, err = writer.Write(rsp)
	if err != nil {
		log.Debug("got write error")
	}
}

func setDefaultHeaders(w http.ResponseWriter) {
	for k, v := range DefaultHeaders {
		w.Header().Set(k, v)
	}
}

func UnmarshallError(handlerName string, w http.ResponseWriter) {
	log.WithFields(log.Fields{
		"handlerName": handlerName,
	}).Debug("UnmarshallError")

	ResponseWithError(w, http.StatusBadRequest, "unmarshall error")
}

func MarshallError(handlerName string, w http.ResponseWriter) {
	log.WithFields(log.Fields{
		"handlerName": handlerName,
	}).Debug("MarshallError")

	ResponseWithError(w, http.StatusBadRequest, "marshall error")
}

func AlreadyExist(handlerName string, w http.ResponseWriter) {
	log.WithFields(log.Fields{
		"handlerName": handlerName,
	}).Debug("AlreadyExist")

	ResponseWithError(w, http.StatusConflict, "entity already exists")
}

func GenerateId() string {
	return uuid.NewString()
}
