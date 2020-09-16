package handlers

import (
	"net/http"
	"photo_album/helpers"
)

func HealthHandler(writer http.ResponseWriter, request *http.Request) {
	response := &helpers.Response{
		Status:  "success",
		Message: "Hi buddy I'm up",
	}
	response.ToResponse(writer)
}

func ReadinessHandler(writer http.ResponseWriter, request *http.Request) {
	response := &helpers.Response{
		Status:  "success",
		Message: "Hi buddy I'm ready",
	}
	response.ToResponse(writer)
}
