package databaseSQL

import (
	"creator/models"
	"database/sql"
	"log"
)

// ART SERVICES
func CreateArt(db *sql.DB, artName string) error {
	_, err := db.Exec(`INSERT INTO arts (art_name) VALUES (?)`, artName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func FindArt(db *sql.DB, artName string) (*models.Art, error) {
	art := &models.Art{}
	err := db.QueryRow(`SELECT arts.id FROM arts WHERE arts.art_name = ?`, artName).Scan(&art.ID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return art, nil
}

func AssignedArtToArtist(db *sql.DB, art *models.Art, artist *models.Artist) error {
	// pass data to table artist-art
	_, err := db.Exec(`INSERT INTO artist_art VALUES (?,?)`, artist.ID, art.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func DeleteArt(db *sql.DB, art *models.Art) error {
	_, err := db.Exec(`DELETE FROM arts WHERE id = ? `, art.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// ARTIST SERVICES
func CreateArtist(db *sql.DB, artistName string) error {
	_, err := db.Exec(`INSERT INTO artists (artist_name) VALUES (?)`, artistName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func FindArtist(db *sql.DB, artistName string) (*models.Artist, error) {
	artist := &models.Artist{}
	err := db.QueryRow(`SELECT artists.id FROM artists WHERE artists.artist_name = ?`, artistName).Scan(&artist.ID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return artist, nil
}

func RegisterArtistToGallery(db *sql.DB, artist *models.Artist, gallery *models.Gallery) error {
	_, err := db.Exec(`INSERT INTO artist_gallery VALUES (?,?)`, artist.ID, gallery.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// GALLERY SERVICES
func CreateGallery(db *sql.DB, g *models.Gallery) {
	_, err := db.Exec(`INSERT INTO galleries (gallery_name) VALUES (?)`, g.Name)
	if err != nil {
		log.Fatal(err)

	}
}

func FindGallery(db *sql.DB, galleryName string) (*models.Gallery, error) {
	g := &models.Gallery{}
	err := db.QueryRow(`SELECT galleries.id FROM galleries WHERE galleries.gallery_name = ?`, galleryName).Scan(&g.ID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return g, nil
}

func UpdateGallery(db *sql.DB, g *models.Gallery, newGalleryName string) error {
	_, err := db.Exec(`UPDATE galleries SET gallery_name = ? WHERE id = ?`, newGalleryName, g.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func DeleteArtist(db *sql.DB, artist *models.Artist, gallery *models.Gallery) error {
	_, err := db.Exec(`DELETE FROM artist_gallery WHERE artist_gallery.artist_id = ? and artist_gallery.gallery_id = ?`, artist.ID, gallery.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
