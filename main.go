//  Documentation for Photo Album MicroService
//
//	RestfulAPI for implementing photo album
//
//  Schemes: http
//  Host: localhost:9090
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//  swagger:meta
package main

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"photo_album/handlers"
	"photo_album/models"
	"time"
)

// TODO-> Decouple image from album
func main() {
	models.DbConnect()
	l := log.New(os.Stdout, "album-api", log.LstdFlags)
	albumHandler := handlers.NewAlbum(l)
	imageHandler := handlers.NewImage(l)
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/album", albumHandler.GetAlbums)
	getRouter.HandleFunc("/album/{album_id:[0-9]+}/image", imageHandler.GetImages)
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/album", albumHandler.AddAlbum)
	postRouter.HandleFunc("/album/{album_id:[0-9]+}/image", imageHandler.AddImage)
	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/album/{album_id:[0-9]+}", albumHandler.DeleteAlbum)
	deleteRouter.HandleFunc("/album/{album_id:[0-9]+}/image/{image_id:[0-9]+}", imageHandler.DeleteImage)
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	getRouter.HandleFunc("/health", handlers.HealthHandler)
	getRouter.HandleFunc("/readiness", handlers.HealthHandler)
	server := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println(sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)

}
