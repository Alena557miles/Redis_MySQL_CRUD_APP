package databaseRedis

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
func CreateArt(rdb *redis.Client, art *models.Art) error {
	bytes, _ := json.Marshal(art)
	err := rdb.Set(ctx, art.Name, bytes, 60*time.Second).Err()
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func FindArt(rdb *redis.Client, artName string) *models.Art {
	artString, err := rdb.Get(ctx, artName).Result()
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

func DeleteArt(rdb *redis.Client, artName string) error {
	_, err := rdb.Get(ctx, artName).Result()
	if err == redis.Nil {
		return nil
	} else {
		err := rdb.Del(ctx, artName).Err()
		if err != nil {
			panic(err)
			return err
		}
		return nil
	}
}

// ARTIST SERVICES
func CreateArtist(rdb *redis.Client, artist *models.Artist) error {
	bytes, _ := json.Marshal(artist)
	err := rdb.Set(ctx, artist.Name, bytes, 60*time.Second).Err()
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func FindArtist(rdb *redis.Client, artistName string) *models.Artist {
	artistString, err := rdb.Get(ctx, artistName).Result()
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
func CreateGallery(rdb *redis.Client, gallery *models.Gallery) error {
	bytes, _ := json.Marshal(gallery)
	err := rdb.Set(ctx, gallery.Name, bytes, 60*time.Second).Err()
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func FindGallery(rdb *redis.Client, galleryName string) *models.Gallery {
	galleryString, err := rdb.Get(ctx, galleryName).Result()
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

func UpdateGallery(rdb *redis.Client, gallery *models.Gallery, newGalleryName string) error {
	log.Println(gallery)
	log.Println(newGalleryName)
	gallery.Name = newGalleryName

	bytes, _ := json.Marshal(gallery)
	log.Println(gallery)
	err := rdb.Set(ctx, newGalleryName, bytes, 60*time.Second).Err()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
