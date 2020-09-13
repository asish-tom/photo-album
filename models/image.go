// As no specification given regarding storage, it's just slices.
// NextID is jus

package models

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"io"
)

// TODO-> Added validations for file
type Image struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
	AlbumID     int    `json:"-"`
}

type Images []*Image

func (i *Image) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}

func (i *Images) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

func (i *Image) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

func AddImage(i *Image, albumId int) (id int, err error) {
	i.Id = getNextImageId()
	al, err := GetAlbumById(albumId)
	if err != nil {
		return
	}
	id = i.Id
	i.AlbumID = al.Id
	imageList = append(imageList, i)
	return
}

func GetImage() Images {
	return imageList
}

func GetImageById(imageId int) (image Image) {
	searchIndex := -1
	for i, each := range imageList {
		if each.Id == imageId {
			searchIndex = i
		}
	}
	if searchIndex > -1 {
		image = *imageList[searchIndex]
	}
	return
}

func DeleteImageById(imageId int) (err error) {
	searchIndex := -1
	for i, each := range imageList {
		if each.Id == imageId {
			searchIndex = i
		}
	}
	if searchIndex == -1 {
		err = errors.New("album-not-found")
		return
	}
	for i := searchIndex; i < len(imageList)-1; i++ {
		imageList[i] = imageList[i+1]
	}
	imageList[len(imageList)-1] = nil
	imageList = imageList[:len(imageList)-1]
	return
}

func getNextImageId() int {
	id := 0
	if len(imageList) > 0 {
		id = imageList[len(imageList)-1].Id
	}
	return id + 1
}

var imageList = []*Image{
	&Image{
		Id:          1,
		Name:        "Image-1",
		AlbumID:     1,
		Description: "Sample Image",
	},
}

func GetImagesByAlbumID(albumId int) (imgList Images) {
	for _, each := range imageList {
		if each.AlbumID == albumId {
			imgList = append(imgList, each)
		}
	}
	return
}
