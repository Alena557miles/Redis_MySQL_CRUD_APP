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
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	artist := &models.Artist{Name: artistName, OnGallery: false}

	err := cache.CreateArtist(artist)
	if err != nil {
		panic(err)
	}

	err = databaseSQL.CreateArtist(artistName)
	if err != nil {
		log.Fatal(err)
	}
	responses.ResponseCreate("Artist", artistName, rw)
}

func (ac *ArtistController) ArtistRegistration(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	var galleryName string = vars["gallery"]

	artist := cache.FindArtist(artistName)
	gallery := cache.FindGallery(galleryName)
	if artist != nil {
		if gallery != nil {
			err := databaseSQL.RegisterArtistToGallery(artist, gallery)
			if err != nil {
				panic(err)
			}
			return
		}
		gallery, err := databaseSQL.FindGallery(galleryName)
		if err != nil {
			panic(err)
		}
		err = databaseSQL.RegisterArtistToGallery(artist, gallery)
		if err != nil {
			log.Fatal(err)
		}
	} else if artist == nil {
		artist, err := databaseSQL.FindArtist(artistName)
		if err != nil {
			panic(err)
		}
		if gallery != nil {
			err = databaseSQL.RegisterArtistToGallery(artist, gallery)
			if err != nil {
				panic(err)
			}
			return
		}
		gallery, err := databaseSQL.FindGallery(galleryName)
		if err != nil {
			panic(err)
		}
		err = databaseSQL.RegisterArtistToGallery(artist, gallery)
		if err != nil {
			panic(err)
		}
	}
	responses.ResponseAction("Artist", artistName, "Gallery", galleryName, "registered", rw)
}

func (ac *ArtistController) DeleteAll(rw http.ResponseWriter, r *http.Request) {

	db, err := databaseSQL.ConnectSQL()
	if err != nil {
		log.Fatalf("SQL DB Connection Failed")
		return
	}
	defer db.Close()
	databaseSQL.PingDB(db)
	err = databaseSQL.DeleteAllArtists()
	if err != nil {
		log.Fatal(err)
	}
	responses.ResponseAction("Artists", "", "", "", "deleteall", rw)
}
