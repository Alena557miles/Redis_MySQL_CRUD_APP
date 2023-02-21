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

	// DELETE ALL ARTS
	// localhost:8080/deleteallarts/
	ac.router.HandleFunc("/deleteallarts/", ac.DeleteAll)

}

func (ac *ArtController) ArtCreation(rw http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var artName = vars["art"]
	art := &models.Art{Name: artName}

	err := cache.CreateArt(art)
	if err != nil {
		log.Println(err)
	}

	err1 := databaseSQL.CreateArt(artName)
	if err1 != nil {
		log.Println(err)
		responses.ResponseError(`Failed to create on DB`, err1, rw)
	}
	responses.ResponseCreate("Art", artName, rw)
}

func (ac *ArtController) AssignArt(rw http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var artistName = vars["artist"]
	var artName = vars["art"]

	art, err := ac.FindValue(artName, rw)
	if err != nil {
		log.Println(err)
		responses.ResponseError(`Failed to assign art to artist on DB`, err, rw)
	}
	artistC := &ArtistController{}
	artist, err1 := artistC.FindValue(artistName, rw)
	if err1 != nil {
		log.Println(err1)
		responses.ResponseError(`Failed to assign art to artist on DB`, err, rw)
	}
	err2 := databaseSQL.AssignedArtToArtist(art, artist)
	if err2 != nil {
		log.Println(err2)
		responses.ResponseError(`Failed to assign art to artist on DB`, err2, rw)
	}
	responses.ResponseAction("Art", artName, "Artist", artistName, "assigned", rw)
}

func (ac *ArtController) ArtDeletion(rw http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var artName = vars["art"]

	err := cache.DeleteArt(artName)
	if err != nil {
		log.Println(err)
	}

	art, err := databaseSQL.FindArt(artName)
	if err != nil {
		log.Println(err)
		responses.ResponseError(`Failed to find on DB`, err, rw)
	}
	err1 := databaseSQL.DeleteArt(art)
	if err1 != nil {
		log.Println(err1)
		responses.ResponseError(`Failed to delete from DB`, err, rw)
	}
	responses.ResponseAction("Art", artName, "", "", "deleted", rw)
}

func (ac *ArtController) DeleteAll(rw http.ResponseWriter, _ *http.Request) {
	err := databaseSQL.DeleteAllArts()
	if err != nil {
		log.Println(err)
		responses.ResponseError(`Failed to delete all arts from DB`, err, rw)
	}
	responses.ResponseAction("Arts", "", "", "", "deleteall", rw)
}

func (ac *ArtController) FindValue(name string, rw http.ResponseWriter) (*models.Art, error) {
	obj := cache.FindArt(name)
	if obj != nil {
		return obj, nil
	} else {
		obj, err := databaseSQL.FindArt(name)
		if err != nil {
			log.Println(err)
			responses.ResponseError(`Failed to find art on DB`, err, rw)
			return nil, err
		}
		return obj, nil
	}
}
