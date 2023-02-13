package controllers

import (
	"creator/databaseRedis"
	"creator/databaseSQL"
	"creator/models"
	"creator/responses"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
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
}

func (ac *ArtistController) Registration(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	artist := &models.Artist{Name: artistName, OnGallery: false}

	rdb := redis.NewClient(databaseRedis.Opt)
	err := databaseRedis.CreateArtist(rdb, artist)
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

	err = databaseSQL.CreateArtist(db, artistName)
	if err != nil {
		log.Fatal(err)
	}
	responses.ResponseCreate("Artist", artistName, rw)
}

func (ac *ArtistController) ArtistRegistration(rw http.ResponseWriter, r *http.Request) {
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

	rdb := redis.NewClient(databaseRedis.Opt)

	artist := databaseRedis.FindArtist(rdb, artistName)
	gallery := databaseRedis.FindGallery(rdb, galleryName)
	if artist != nil {
		if gallery != nil {
			err = databaseSQL.RegisterArtistToGallery(db, artist, gallery)
			if err != nil {
				panic(err)
			}
			return
		}
		gallery, err := databaseSQL.FindGallery(db, galleryName)
		if err != nil {
			panic(err)
		}
		err = databaseSQL.RegisterArtistToGallery(db, artist, gallery)
		if err != nil {
			log.Fatal(err)
		}
	} else if artist == nil {
		artist, err := databaseSQL.FindArtist(db, artistName)
		if err != nil {
			panic(err)
		}
		if gallery != nil {
			err = databaseSQL.RegisterArtistToGallery(db, artist, gallery)
			if err != nil {
				panic(err)
			}
			return
		}
		gallery, err := databaseSQL.FindGallery(db, galleryName)
		if err != nil {
			panic(err)
		}
		err = databaseSQL.RegisterArtistToGallery(db, artist, gallery)
		if err != nil {
			panic(err)
		}
	}
	responses.ResponseAction("Artist", artistName, "Gallery", galleryName, "registered", rw)
}
