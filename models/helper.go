package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type GormClient struct {
	crDB *gorm.DB
}

var err error
var host string
var gc GormClient

func DbConnect() {
	var connectionString = "photo-album:photo-album@tcp(" + host + ")/photo-album?charset=utf8mb4&parseTime=True&loc=Local"
	gc.crDB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}
	// Migrate the schema
	gc.crDB.AutoMigrate(&Image{}, &Album{})
	fmt.Print("Connection success")
}

func init() {
	host = "localhost"
	if os.Getenv("DB_HOST") != "" {
		host = os.Getenv("DB_HOST")
	}
	fmt.Println(host)
}
