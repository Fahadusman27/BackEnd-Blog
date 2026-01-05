package model

import "time"

type Post struct {
	ID             int        `json:"id"`
	UserID         []Users    `json:"Users_ID"`
	KategoriID     []Kategori `json:"Kategori_ID"`
	Title          string     `json:"Title"`
	Slug           string     `json:"Slug"`
	Content        string     `json:"COntent"`
	Featured_Image string     `json:"featured_image"`
	Status         string     `json:"Status"`
	CreatedAt      time.Time  `json:"CreatedAt"`
	UpdateAt       time.Time  `json:"UpdateAt"`
}
