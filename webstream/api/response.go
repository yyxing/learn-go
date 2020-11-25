package main

import (
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

type HttpStatusCode int

const (
	SUCCESS           HttpStatusCode = http.StatusOK
	HttpNotFound      HttpStatusCode = http.StatusNotFound
	HttpForbidden     HttpStatusCode = http.StatusForbidden
	HttpUnauthorized  HttpStatusCode = http.StatusUnauthorized
	HttpBadRequest    HttpStatusCode = http.StatusBadRequest
	HttpInternalError HttpStatusCode = http.StatusInternalServerError
)

var (
	json                        = jsoniter.ConfigCompatibleWithStandardLibrary
	RequestBodyParseFailedError = HttpResponse{Code: HttpBadRequest, Message: "Request body is not correct"}
	NotAuthUserError            = HttpResponse{Code: HttpUnauthorized, Message: "User authentication failed."}
	ServerError                 = HttpResponse{Code: HttpInternalError, Message: "Server abnormality, please try again later"}
)

type HttpResponse struct {
	Code    HttpStatusCode
	Message string
	Data    interface{}
}

func handleError(response HttpResponse, w http.ResponseWriter) {
	sendResponse(w, response)
}
func sendResponse(w http.ResponseWriter, response HttpResponse) {
	w.WriteHeader(int(response.Code))
	resStr, _ := json.Marshal(&response)
	_, err := io.WriteString(w, string(resStr))
	if err != nil {
		panic(err)
	}
}
func Success(message string, data interface{}) HttpResponse {
	return HttpResponse{
		Code:    SUCCESS,
		Message: message,
		Data:    data,
	}
}
func Fail(code HttpStatusCode, message string, data interface{}) HttpResponse {
	return HttpResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
