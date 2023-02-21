package controllers

import (
	"creator/cache"
	"creator/databaseSQL"
	"creator/models"
	"creator/responses"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type GalleryController struct {
	Galleries []*models.Gallery
	router    *mux.Router
}

func (gc *GalleryController) RegisterRouter(r *mux.Router) {
	gc.router = r
}

func (gc *GalleryController) RegisterActions() {
	// CREATE GALLERY
	// localhost:8080/creategallery/Tokio
	gc.router.HandleFunc("/creategallery/{gallery}", gc.GalleryCreation)

	// DELETE AN ARTIST FROM GALLERY
	// localhost:8080/artist/delete/Fillip/Tokio
	gc.router.HandleFunc("/artist/delete/{artist}/{gallery}", gc.RemoveArtistFromGal)

	// RENAME GALLERY
	// localhost:8080/renamegallery/Tokio/JapaneTreasure
	gc.router.HandleFunc("/renamegallery/{gallery}/{newgallery}", gc.GalleryUpdate)

	// DELETE ALL GALLERIES
	// localhost:8080/deleteallgalleries/
	gc.router.HandleFunc("/deleteallgalleries", gc.DeleteAll)
}

func (gc *GalleryController) GalleryCreation(rw http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var galleryName = vars["gallery"]
	gallery := &models.Gallery{Name: galleryName}

	err1 := databaseSQL.CreateGallery(gallery)
	if err1 != nil {
		log.Println(err1)
		responses.ResponseError(`Failed to create on DB`, err1, rw)
	}
	err := cache.CreateGallery(gallery)
	if err != nil {
		log.Println(err)
		responses.ResponseError(`Failed to create cache`, err, rw)
	}
	responses.ResponseCreate("Gallery", galleryName, rw)
}

func (gc *GalleryController) RemoveArtistFromGal(rw http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var artistName = vars["artist"]
	var galleryName = vars["gallery"]
	gallery, err := gc.FindValue(galleryName, rw)
	if err != nil {
		log.Println(err)
		responses.ResponseError(`Failed to find on DB `, err, rw)
	}
	artistC := &ArtistController{}
	artist, err1 := artistC.FindValue(artistName, rw)
	if err1 != nil {
		log.Println(err1)
		responses.ResponseError(`Failed to assign art to artist on DB`, err, rw)
	}
	err2 := databaseSQL.DeleteArtist(artist, gallery)
	if err2 != nil {
		log.Println(err2)
		responses.ResponseError(`Failed to delete from DB `, err1, rw)
	}
	responses.ResponseAction("Artist", artistName, "Gallery", galleryName, "deleted", rw)
}

func (gc *GalleryController) GalleryUpdate(rw http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var galleryName = vars["gallery"]
	var newGalleryName = vars["newgallery"]

	g, err := databaseSQL.FindGallery(galleryName)
	if err != nil {
		log.Println(err)
		responses.ResponseError(`Failed to find on DB `, err, rw)
	}
	err1 := databaseSQL.UpdateGallery(g, newGalleryName)
	if err1 != nil {
		log.Println(err1)
		responses.ResponseError(`Failed to update on DB `, err1, rw)
	}
	err2 := cache.UpdateGallery(g, newGalleryName)
	if err2 != nil {
		log.Println(err2)
		responses.ResponseError(`Failed to update on DB `, err2, rw)
	}

	responses.ResponseAction("Gallery", galleryName, "New gallery name", newGalleryName, "update", rw)
}

func (gc *GalleryController) DeleteAll(rw http.ResponseWriter, _ *http.Request) {
	err := databaseSQL.DeleteAllGalleries()
	if err != nil {
		responses.ResponseError(`Failed to delete from DB`, err, rw)
	}
	responses.ResponseAction("Galleries", "", "", "", "deleteall", rw)
}

func (gc *GalleryController) FindValue(name string, rw http.ResponseWriter) (*models.Gallery, error) {
	obj := cache.FindGallery(name)
	if obj != nil {
		return obj, nil
	} else {
		obj, err := databaseSQL.FindGallery(name)
		if err != nil {
			log.Println(err)
			responses.ResponseError(`Failed to find artis on DB`, err, rw)
			return nil, err
		}
		return obj, nil
	}
}
