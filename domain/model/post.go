package model

import "time"

type Post struct {
	ID             int       `json:"id"`
	UserID         int       `json:"Users_ID"`
	KategoriID     int       `json:"Kategori_ID"`
	TitleID        int       `json:"Title_ID"`
	LikeID         int       `json:"LikeID"`
	ShareID        int       `json:"ShareID"`
	Slug           string    `json:"Slug"`
	Content        string    `json:"Content"`
	Featured_Image string    `json:"featured_image"`
	Status         string    `json:"Status"`
	CreatedAt      time.Time `json:"CreatedAt"`
	UpdateAt       time.Time `json:"UpdateAt"`
}
