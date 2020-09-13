package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"photo_album/helpers"
	"photo_album/models"
	"strconv"
)

type Images struct {
	l *log.Logger
}

// swagger:route GET /album/{album_id}/image image listAllImagesOfSelectedImages
// Returns a list of images in the album
func (i *Images) GetImages(writer http.ResponseWriter, request *http.Request) {
	i.l.Println("Handle Get image")
	vars := mux.Vars(request)
	albumId, err := strconv.Atoi(vars["album_id"])
	if err != nil {
		http.Error(writer, "Unable to identify album", http.StatusBadRequest)
	}
	li := models.GetImagesByAlbumID(albumId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	err = li.ToJSON(writer)
	if err != nil {
		http.Error(writer, "Unable to pack", http.StatusInternalServerError)
	}

}

// swagger:route POST /album/{album_id}/image image addImageToAlbum
// Returns a 200 in case of success.
func (i *Images) AddImage(writer http.ResponseWriter, request *http.Request) {
	i.l.Println("Handle post image")
	vars := mux.Vars(request)
	img := &models.Image{}
	albumId, err := strconv.Atoi(vars["album_id"])
	if err != nil {
		http.Error(writer, "Unable to identify album", http.StatusBadRequest)
	}
	err = img.FromJSON(request.Body)
	if err != nil {
		http.Error(writer, "Unable to unmarshal json", http.StatusBadRequest)
	}
	id, err := models.AddImage(img, albumId)
	data := map[string]interface{}{
		"image_id": id,
		"event":    "Added Image",
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	defer helpers.PublishToKafka("albumNotification", data)
}

// swagger:route Delete /album/{album_id}/image/{image_id} image deleteImageFromAlbum
// Returns a 200 in case of success.
func (a *Images) DeleteImage(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["image_id"])
	if err != nil {
		http.Error(writer, "unable-to-extract-id", http.StatusBadRequest)
	}
	err = models.DeleteImageById(id)
	if err != nil {
		http.Error(writer, "unable-to-find-image", http.StatusNotFound)
	}
	data := map[string]interface{}{
		"image_id": id,
		"event":    "Deleted Image",
	}
	defer helpers.PublishToKafka("albumNotification", data)
	return
}

func NewImage(l *log.Logger) *Images {
	return &Images{l}
}
