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
}

func (gc *GalleryController) GalleryCreation(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var galleryName string = vars["gallery"]
	gallery := &models.Gallery{Name: galleryName}

	err := cache.CreateGallery(gallery)
	if err != nil {
		panic(err)
	}

	db, err := databaseSQL.ConnectSQL()
	if err != nil {
		log.Fatalf("SQL DB Connection Failed")
		return
	}
	defer db.Close()
	databaseSQL.PingDB(db)
	databaseSQL.CreateGallery(db, gallery)
	responses.ResponseCreate("Gallery", galleryName, rw)
}

func (gc *GalleryController) RemoveArtistFromGal(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	var galleryName string = vars["gallery"]

	db, err := databaseSQL.ConnectSQL()
	if err != nil {
		log.Fatalf("SQL DB Connection Failed")
		return
	}
	defer db.Close()
	databaseSQL.PingDB(db)

	artist := cache.FindArtist(artistName)
	gallery := cache.FindGallery(galleryName)
	if artist != nil {
		if gallery != nil {
			err = databaseSQL.DeleteArtist(db, artist, gallery)
			if err != nil {
				panic(err)
			}
			return
		}
		gallery, err := databaseSQL.FindGallery(db, galleryName)
		if err != nil {
			log.Fatal(err)
		}
		err = databaseSQL.DeleteArtist(db, artist, gallery)
		if err != nil {
			panic(err)
		}
	} else if artist == nil {
		// searching data on DB MySQL
		artist, err := databaseSQL.FindArtist(db, artistName)
		if err != nil {
			log.Fatal(err)
		}
		if gallery != nil {
			err = databaseSQL.DeleteArtist(db, artist, gallery)
			if err != nil {
				panic(err)
			}
			return
		}
		gallery, err := databaseSQL.FindGallery(db, galleryName)
		if err != nil {
			log.Fatal(err)
		}
		err = databaseSQL.DeleteArtist(db, artist, gallery)
		if err != nil {
			log.Fatal(err)
		}
	}
	responses.ResponseAction("Artist", artistName, "Gallery", galleryName, "deleted", rw)
}

func (gc *GalleryController) GalleryUpdate(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var galleryName string = vars["gallery"]
	var newGalleryName string = vars["newgallery"]

	db, err := databaseSQL.ConnectSQL()
	if err != nil {
		log.Fatalf("SQL DB Connection Failed")
		return
	}
	defer db.Close()
	databaseSQL.PingDB(db)

	g, err := databaseSQL.FindGallery(db, galleryName)
	if err != nil {
		log.Fatal(err)
	}
	err = databaseSQL.UpdateGallery(db, g, newGalleryName)
	if err != nil {
		log.Fatal(err)
	}
	err = cache.UpdateGallery(g, newGalleryName)
	if err != nil {
		panic(err)
	}

	responses.ResponseAction("Gallery", galleryName, "New gallery name", newGalleryName, "update", rw)
}
