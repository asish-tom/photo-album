// As no specification given regarding storage, it's just slices.
// NextID is jus

package models

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"io"
)

// TODO-> Added validations for file
type Image struct {
	gorm.Model  `gorm:"embedded"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
	AlbumID     uint   `json:"-"`
	ID          uint   `json:"id"`
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

func AddImage(i *Image, albumId uint) (id uint, err error) {
	al, err := GetAlbumById(albumId)
	if err != nil {
		err = errors.New("Album dose not exists")
		return
	}
	id = i.ID
	i.AlbumID = al.ID
	txn := gc.crDB.Create(&i)
	if txn.Error != nil {
		return 0, err
	}
	id = i.ID
	return
}

func DeleteImageById(imageId uint) (err error) {
	var i = &Image{}
	txn := gc.crDB.Unscoped().Delete(&i, imageId)
	if txn.Error != nil {
		err = txn.Error
	} else if txn.RowsAffected == 0 {
		err = errors.New("Image dose not exists")
	}
	return
}

func GetImagesByAlbumID(albumId int) (imgList Images) {
	_ = gc.crDB.Where("album_id= ? ", albumId).Find(&imgList)
	return
}
