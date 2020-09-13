package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"photo_album/models"
	"strconv"
)

type Albums struct {
	l *log.Logger
}

// swagger:route GET /album albums listAlbums
// Returns a list of all albums in the system
func (a *Albums) GetAlbums(writer http.ResponseWriter, request *http.Request) {
	a.l.Println("Handle GET Albums")
	la := models.GetAlbums()
	err := la.ToJSON(writer)
	if err != nil {
		http.Error(writer, "Unable to pack", http.StatusInternalServerError)
	}
}

// swagger:route POST /album album saveAlbum
// Returns a 200 in case of success
func (a *Albums) AddAlbum(writer http.ResponseWriter, request *http.Request) {
	a.l.Println("Handle POST Albums")
	alb := &models.Album{}
	err := alb.FromJSON(request.Body)
	if err != nil {
		http.Error(writer, "Unable to unmarshal json", http.StatusBadRequest)
	}
	err = alb.Validate()
	if err != nil {
		http.Error(writer, "json-validation-failed", http.StatusBadRequest)
	}
	a.l.Printf("AlbumID: %#v", alb)
	models.AddAlbum(alb)
}

// swagger:route DELETE /album/{id} album deleteAlbum
// Returns a 200 in case of success
func (a *Albums) DeleteAlbum(writer http.ResponseWriter, request *http.Request) {
	a.l.Println("Handle Delete Albums")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["album_id"])
	if err != nil {
		http.Error(writer, "unable-to-extract-id", http.StatusBadRequest)
	}
	err = models.DeleteAlbumById(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
	}
	return
}

func NewAlbum(l *log.Logger) *Albums {
	return &Albums{l}
}
