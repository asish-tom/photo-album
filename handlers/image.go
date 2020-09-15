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
// responses:
//	200:Response
func (i *Images) GetImages(writer http.ResponseWriter, request *http.Request) {
	i.l.Println("Handle GET album images")
	response := &helpers.Response{
		Status:  "failed",
		Message: "Unknown error",
	}
	vars := mux.Vars(request)
	albumId, err := strconv.Atoi(vars["album_id"])
	if err != nil {
		response.Message = "Unable to identify album"
		response.ToResponse(writer)
		return
	}
	li := models.GetImagesByAlbumID(albumId)
	if len(li) == 0 {
		response.Message = "Retrieved albums is  empty"
	} else {
		response.Message = "Retrieved album images successfully"
	}
	response.Status = "success"
	response.Model = li
	response.ToResponse(writer)
	return

}

// swagger:route POST /album/{album_id}/image image addImageToAlbum
// Returns a 200 in case of success.
// responses:
//	200:Response
func (i *Images) AddImage(writer http.ResponseWriter, request *http.Request) {
	i.l.Println("Handle POST album image")
	response := &helpers.Response{
		Status:  "failed",
		Message: "Unknown error",
	}
	vars := mux.Vars(request)
	img := &models.Image{}
	albumId, err := strconv.Atoi(vars["album_id"])
	if err != nil {
		response.Message = "Unable to identify album"
		response.ToResponse(writer)
		return
	}
	err = img.FromJSON(request.Body)
	if err != nil {
		response.Message = "Unable to unmarshal json"
		response.ToResponse(writer)
		return
	}
	err = img.Validate()
	if err != nil {
		response.Message = "Unable to validate json"
		response.ToResponse(writer)
		return
	}

	id, err := models.AddImage(img, uint(albumId))
	if err != nil {
		response.Message = err.Error()
		response.ToResponse(writer)
		return
	}
	data := map[string]interface{}{
		"image_id": id,
		"event":    "Added Image",
	}
	response.Message = "Added image successfully"
	response.Status = "success"
	response.Model = img
	response.ToResponse(writer)
	defer helpers.PublishToKafka("albumNotification", data)
}

// swagger:route Delete /album/{album_id}/image/{image_id} image deleteImageFromAlbum
// Returns a 200 in case of success.
// responses:
//	200:Response
func (i *Images) DeleteImage(writer http.ResponseWriter, request *http.Request) {
	response := &helpers.Response{
		Status:  "failed",
		Message: "Unknown error",
	}
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["image_id"])
	if err != nil {
		response.Message = "Unable to extract image id"
		response.ToResponse(writer)
		return
	}
	err = models.DeleteImageById(uint(id))
	if err != nil {
		response.Message = "Unable to find the image"
		response.ToResponse(writer)
		return
	}
	data := map[string]interface{}{
		"image_id": id,
		"event":    "Deleted Image",
	}
	defer helpers.PublishToKafka("albumNotification", data)
	response.Message = "Deleted image successfully"
	response.Status = "success"
	response.ToResponse(writer)
	defer helpers.PublishToKafka("albumNotification", data)
	return
}

func NewImage(l *log.Logger) *Images {
	return &Images{l}
}
