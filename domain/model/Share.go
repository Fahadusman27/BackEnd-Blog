package model

import "time"

type Share struct {
	ID       int       `json:"ID"`
	UserID   int       `json:"UserID"`
	PostID   int       `json:"PostID"`
	Platform string    `json:"Platform"`
	CreateAt time.Time `json:"CreateAt"`
}
