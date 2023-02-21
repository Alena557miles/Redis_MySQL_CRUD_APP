package databaseSQL

import (
	"creator/models"
	"log"
)

// ART SERVICES
func CreateArt(artName string) error {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return err
	}
	_, err = db.Exec(`INSERT INTO arts (art_name) VALUES (?)`, artName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindArt(artName string) (*models.Art, error) {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return nil, err
	}
	art := &models.Art{}
	err = db.QueryRow(`SELECT art_id FROM arts WHERE arts.art_name = ?`, artName).Scan(&art.ID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println(art)
	return art, nil
}

func AssignedArtToArtist(art *models.Art, artist *models.Artist) error {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return err
	}
	_, err = db.Exec(`INSERT INTO artist_art VALUES (?,?)`, artist.ID, art.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteArt(art *models.Art) error {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return err
	}
	_, err = db.Exec(`DELETE FROM arts WHERE id = ? `, art.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteAllArts() error {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return err
	}
	_, err = db.Exec(`DELETE FROM arts`)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ARTIST SERVICES
func CreateArtist(artistName string) error {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return err
	}
	_, err = db.Exec(`INSERT INTO artists (artist_name) VALUES (?)`, artistName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindArtist(artistName string) (*models.Artist, error) {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return nil, err
	}
	artist := &models.Artist{}
	err = db.QueryRow(`SELECT artist_id FROM artists WHERE artists.artist_name = ?`, artistName).Scan(&artist.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return artist, nil
}

func RegisterArtistToGallery(artist *models.Artist, gallery *models.Gallery) error {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return err
	}

	_, err = db.Exec(`INSERT INTO artist_gallery VALUES (?,?)`, artist.ID, gallery.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteAllArtists() error {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return err
	}
	_, err = db.Exec(`DELETE FROM artists`)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// GALLERY SERVICES
func CreateGallery(g *models.Gallery) error {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return err
	}
	_, err = db.Exec(`INSERT INTO galleries (gallery_name) VALUES (?)`, g.Name)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func FindGallery(galleryName string) (*models.Gallery, error) {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return nil, err
	}

	g := &models.Gallery{}
	err = db.QueryRow(`SELECT gallery_id FROM galleries WHERE galleries.gallery_name = ?`, galleryName).Scan(&g.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return g, nil
}

func UpdateGallery(g *models.Gallery, newGalleryName string) error {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return err
	}
	_, err = db.Exec(`UPDATE galleries SET gallery_name = ? WHERE id = ?`, newGalleryName, g.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteArtist(artist *models.Artist, gallery *models.Gallery) error {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return err
	}
	_, err = db.Exec(`DELETE FROM artist_gallery WHERE artist_gallery.artist_id = ? and artist_gallery.gallery_id = ?`, artist.ID, gallery.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteAllGalleries() error {
	db, err := GetInstance()
	if err != nil {
		log.Println("SQL DB Connection Failed")
		return err
	}
	_, err = db.Exec(`DELETE FROM galleries`)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
