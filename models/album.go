// As no specification given regarding storage, it's just slices.
package models

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"io"
)

type Album struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type Albums []*Album

func (a *Albums) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *Album) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

func (a *Album) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}

func GetAlbums() Albums {
	var queryList []*Album
	for _, each := range albumList {
		queryList = append(queryList, each)
	}
	return queryList
}

func AddAlbum(a *Album) int {
	a.Id = getNextAlbumId()
	albumList = append(albumList, a)
	return a.Id
}

func DeleteAlbumById(id int) (err error) {
	searchIndex := -1
	// searching for the album
	for i, each := range albumList {
		if each.Id == id {
			searchIndex = i
		}
	}
	if searchIndex == -1 {
		err = errors.New("album-not-found")
		return
	}
	// searching if album contains image
	for _, each := range imageList {
		if each.AlbumID == albumList[searchIndex].Id {
			err = errors.New("images are associated with album")
			return
		}
	}
	// deleting album
	for i := searchIndex; i < len(albumList); i++ {
		albumList[i] = albumList[i+1]
	}
	albumList[len(albumList)-1] = nil
	albumList = albumList[:len(albumList)-1]
	return
}

func GetAlbumById(id int) (album Album, err error) {
	searchIndex := -1
	for i, each := range albumList {
		if each.Id == id {
			searchIndex = i
		}
	}
	if searchIndex == -1 {
		err = errors.New("album-not-found")
		return
	}
	album = *albumList[searchIndex]
	return
}

func getNextAlbumId() int {
	id := 0
	if len(albumList) > 0 {
		id = albumList[len(albumList)-1].Id
	}
	return id + 1
}

var albumList = []*Album{
	{
		Id:          1,
		Name:        "Album1",
		Description: "Sample album 1",
	},
	{
		Id:          2,
		Name:        "Album2",
		Description: "Sample album2",
	},
}
