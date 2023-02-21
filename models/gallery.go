package models

type Gallery struct {
	ID      string    `json:"gallery_id"`
	Name    string    `json:"gallery_name"`
	Artists []*Artist `json:"gallery_artists"`
}
