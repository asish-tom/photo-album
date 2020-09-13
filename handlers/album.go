package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"photo_album/helpers"
	"photo_album/models"
	"strconv"
)

type Albums struct {
	l *log.Logger
}

// swagger:route GET /album album listAlbums
// Returns a list of all albums in the system
// responses:
//	200:ResponseModel
func (a *Albums) GetAlbums(writer http.ResponseWriter, request *http.Request) {
	a.l.Println("Handle GET Album")
	response := &helpers.ResponseModel{
		Status:  "failed",
		Message: "Unknown error",
	}
	a.l.Println("Handle GET Albums")
	la := models.GetAlbums()
	response.Model = la
	response.Status = "success"
	response.Message = ""
	response.ToResponse(writer)
	return
}

// swagger:route POST /album album saveAlbum
// Returns a 200 in case of success
// responses:
//	200:ResponseModel
func (a *Albums) AddAlbum(writer http.ResponseWriter, request *http.Request) {
	a.l.Println("Handle POST Album")
	response := &helpers.ResponseModel{
		Status:  "failed",
		Message: "Unknown error",
	}
	alb := &models.Album{}
	err := alb.FromJSON(request.Body)
	if err != nil {
		response.Message = "Unable to unmarshal json"
		response.ToResponse(writer)
		return
	}
	err = alb.Validate()
	if err != nil {
		response.Message = "json-validation-failed"
		response.ToResponse(writer)
		return
	}
	id := models.AddAlbum(alb)
	alb.Id = id
	response.Model = alb
	response.Status = "success"
	response.Message = ""
	response.ToResponse(writer)
	return
}

// swagger:route DELETE /album/{id} album deleteAlbum
// Returns a 200 in case of success
// responses:
//	200:ResponseModel
func (a *Albums) DeleteAlbum(writer http.ResponseWriter, request *http.Request) {
	a.l.Println("Handle Delete Album")
	response := &helpers.ResponseModel{
		Status:  "failed",
		Message: "Unknown error",
	}
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["album_id"])
	if err != nil {
		response.Message = "Unable to extract album id"
		response.ToResponse(writer)
		return
	}
	err = models.DeleteAlbumById(id)
	if err != nil {
		response.Message = err.Error()
		response.ToResponse(writer)
		return
	}
	response.Status = "success"
	response.Message = "Album deleted"
	response.ToResponse(writer)
	return
}

func NewAlbum(l *log.Logger) *Albums {
	return &Albums{l}
}
