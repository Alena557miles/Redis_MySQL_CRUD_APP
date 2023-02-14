package controllers

import (
	"context"
	"creator/cache"
	"creator/databaseSQL"
	"creator/models"
	"creator/responses"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var ctx = context.Background()

type ArtController struct {
	arts   []*models.Art
	router *mux.Router
}

func (ac *ArtController) RegisterRouter(r *mux.Router) {
	ac.router = r
}

func (ac *ArtController) RegisterActions() {
	// CREATE AN ART
	// localhost:8080/createart/blackCat
	ac.router.HandleFunc("/createart/{art}", ac.ArtCreation)

	//ASSIGN AN ART TO THE ARTIST (BY NAME)
	// localhost:8080/artist/assign/Fillip/blackCat
	ac.router.HandleFunc("/artist/assign/{artist}/{art}", ac.AssignArt)

	// DELETE AN ART
	// localhost:8080/deleteart/blackCat
	ac.router.HandleFunc("/deleteart/{art}", ac.ArtDeletion)

}

func (ac *ArtController) ArtCreation(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artName string = vars["art"]
	art := &models.Art{Name: artName}

	err := cache.CreateArt(art)
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
	err = databaseSQL.CreateArt(db, artName)
	if err != nil {
		log.Fatal(err)
	}
	responses.ResponseCreate("Art", artName, rw)
}

func (ac *ArtController) AssignArt(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	var artName string = vars["art"]

	db, err := databaseSQL.ConnectSQL()
	if err != nil {
		log.Fatalf("SQL DB Connection Failed")
		return
	}
	defer db.Close()
	databaseSQL.PingDB(db)

	art := cache.FindArt(artName)
	artist := cache.FindArtist(artistName)
	if art != nil {
		if artist != nil {
			err = databaseSQL.AssignedArtToArtist(db, art, artist)
			if err != nil {
				panic(err)
			}
			return
		}
		artist, err := databaseSQL.FindArtist(db, artistName)
		if err != nil {
			panic(err)
		}
		err = databaseSQL.AssignedArtToArtist(db, art, artist)
		if err != nil {
			panic(err)
		}
	} else if art == nil {
		art, err := databaseSQL.FindArt(db, artName)
		if err != nil {
			panic(err)
		}
		if artist != nil {
			err = databaseSQL.AssignedArtToArtist(db, art, artist)
			if err != nil {
				panic(err)
			}
			return
		}
		artist, err := databaseSQL.FindArtist(db, artistName)
		if err != nil {
			panic(err)
		}
		err = databaseSQL.AssignedArtToArtist(db, art, artist)
		if err != nil {
			panic(err)
		}
	}

	responses.ResponseAction("Art", artName, "Artist", artistName, "assigned", rw)
}

func (ac *ArtController) ArtDeletion(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artName string = vars["art"]

	err := cache.DeleteArt(artName)
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

	art, err := databaseSQL.FindArt(db, artName)
	if err != nil {
		panic(err)
	}
	err = databaseSQL.DeleteArt(db, art)
	if err != nil {
		log.Fatal(err)
	}
	responses.ResponseAction("Art", artName, "", "", "deleted", rw)
}
