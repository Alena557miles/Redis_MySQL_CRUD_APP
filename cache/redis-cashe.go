package cache

import (
	"context"
	"creator/models"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var ctx = context.Background()

// ART SERVICES
func CreateArt(art *models.Art) error {
	client := GetInstance()
	bytes, _ := json.Marshal(art)
	err := client.Set(ctx, art.Name, bytes, 60*time.Second).Err()
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func FindArt(artName string) *models.Art {
	client := GetInstance()
	artString, err := client.Get(ctx, artName).Result()
	if err == redis.Nil {
		log.Println("Art does not exist in Redis")
		return nil
	} else if err != nil {
		panic(err)
	} else {
		log.Println("Data about Art from Redis: ", artString)
		artToByte := []byte(artString)
		var art *models.Art
		json.Unmarshal(artToByte, &art)
		return art
	}
}

func DeleteArt(artName string) error {
	client := GetInstance()
	_, err := client.Get(ctx, artName).Result()
	if err == redis.Nil {
		return nil
	} else {
		err := client.Del(ctx, artName).Err()
		if err != nil {
			panic(err)
			return err
		}
		return nil
	}
}

// ARTIST SERVICES
func CreateArtist(artist *models.Artist) error {
	client := GetInstance()
	bytes, _ := json.Marshal(artist)
	err := client.Set(ctx, artist.Name, bytes, 60*time.Second).Err()
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func FindArtist(artistName string) *models.Artist {
	client := GetInstance()
	artistString, err := client.Get(ctx, artistName).Result()
	if err == redis.Nil {
		log.Println("Artist does not exist in Redis")
		return nil
	} else if err != nil {
		panic(err)
	} else {
		log.Println("Data about Artist from Redis: ", artistString)
		artToByte := []byte(artistString)
		var artist *models.Artist
		json.Unmarshal(artToByte, &artist)
		return artist
	}
}

// GALLERY SERVICES
func CreateGallery(gallery *models.Gallery) error {
	client := GetInstance()
	bytes, _ := json.Marshal(gallery)
	err := client.Set(ctx, gallery.Name, bytes, 60*time.Second).Err()
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func FindGallery(galleryName string) *models.Gallery {
	client := GetInstance()
	galleryString, err := client.Get(ctx, galleryName).Result()
	if err == redis.Nil {
		log.Println("Gallery does not exist in Redis")
		return nil
	} else if err != nil {
		panic(err)
	} else {
		log.Println("Data about Gallery from Redis: ", galleryString)
		galleryToByte := []byte(galleryString)
		var gallery *models.Gallery
		json.Unmarshal(galleryToByte, &gallery)
		return gallery
	}
}

func UpdateGallery(gallery *models.Gallery, newGalleryName string) error {
	client := GetInstance()
	log.Println(gallery)
	log.Println(newGalleryName)
	gallery.Name = newGalleryName

	bytes, _ := json.Marshal(gallery)
	log.Println(gallery)
	err := client.Set(ctx, newGalleryName, bytes, 60*time.Second).Err()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
