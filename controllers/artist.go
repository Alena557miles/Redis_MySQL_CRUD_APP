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

type ArtistController struct {
	artists []*models.Artist
	router  *mux.Router
}

func (ac *ArtistController) RegisterRouter(r *mux.Router) {
	ac.router = r
}

func (ac *ArtistController) RegisterActions() {
	// CREATE AN ARTIST
	// localhost:8080/createartist/Fillip
	ac.router.HandleFunc("/createartist/{artist}", ac.Registration)

	//REGISTRATION AN ARTIST ON THE GALLERY
	// localhost:8080/artist/register/Fillip/Tokio
	ac.router.HandleFunc("/artist/register/{artist}/{gallery}", ac.ArtistRegistration)

	//DELETE ALL ARTISTS
	// localhost:8080/deleteallartists/
	ac.router.HandleFunc("/deleteallartists", ac.DeleteAll)
}

func (ac *ArtistController) Registration(rw http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var artistName = vars["artist"]
	artist := &models.Artist{Name: artistName, OnGallery: false}

	err := cache.CreateArtist(artist)
	if err != nil {
		log.Println(err)
	}

	err1 := databaseSQL.CreateArtist(artistName)
	if err1 != nil {
		responses.ResponseError(`Failed to create on DB`, err1, rw)
	}
	responses.ResponseCreate("Artist", artistName, rw)
}

func (ac *ArtistController) ArtistRegistration(rw http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var artistName = vars["artist"]
	var galleryName = vars["gallery"]

	artist, err := ac.FindValue(artistName, rw)
	if err != nil {
		log.Println(err)
		responses.ResponseError(`Failed to find artist on DB`, err, rw)
	}
	galleryC := &GalleryController{}
	gallery, err1 := galleryC.FindValue(galleryName, rw)
	if err1 != nil {
		log.Println(err1)
		responses.ResponseError(`Failed to find gallery on DB`, err, rw)
	}
	err2 := databaseSQL.RegisterArtistToGallery(artist, gallery)
	if err2 != nil {
		log.Println(err2)
		responses.ResponseError(`Failed to register `, err2, rw)
	}
	responses.ResponseAction("Artist", artistName, "Gallery", galleryName, "registered", rw)
}

func (ac *ArtistController) DeleteAll(rw http.ResponseWriter, _ *http.Request) {
	err := databaseSQL.DeleteAllArtists()
	if err != nil {
		responses.ResponseError(`Failed to delete from DB`, err, rw)
	}
	responses.ResponseAction("Artists", "", "", "", "deleteall", rw)
}

func (ac *ArtistController) FindValue(name string, rw http.ResponseWriter) (*models.Artist, error) {
	obj := cache.FindArtist(name)
	if obj != nil {
		return obj, nil
	} else {
		obj, err := databaseSQL.FindArtist(name)
		if err != nil {
			log.Println(err)
			responses.ResponseError(`Failed to find artis on DB`, err, rw)
			return nil, err
		}
		return obj, nil
	}
}
