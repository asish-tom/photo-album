// As no specification given regarding storage, it's just slices.
package models

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"io"
)

type Album struct {
	gorm.Model  `gorm:"embedded"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ID          uint
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

func GetAlbums() (albs Albums) {
	_ = gc.crDB.Find(&albs)
	return
}

func AddAlbum(a *Album) uint {
	_ = gc.crDB.Create(&a)
	return a.ID
}

func DeleteAlbumById(id uint) (err error) {

	images := Images{}
	txn := gc.crDB.Where("album_id= ? ", id).Find(&images)
	if txn.Error != nil {
		err = txn.Error
		return
	}
	if len(images) > 0 {
		err = errors.New("images are associated with album")
	}

	// deleting album
	alb := Album{}
	txn = gc.crDB.Unscoped().Delete(&alb, id)
	if txn.Error != nil {
		err = txn.Error
	} else if txn.RowsAffected == 0 {
		err = errors.New("Album dose not exists")
	}
	print(txn.RowsAffected)
	return
}

func GetAlbumById(id uint) (album Album, err error) {
	alb := Album{
		ID: id,
	}
	txn := gc.crDB.First(&alb)
	if txn.Error != nil {
		err = txn.Error
	}
	return
}
