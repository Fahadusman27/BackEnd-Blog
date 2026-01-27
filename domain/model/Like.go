package model

import "time"

type Like struct {
	ID        int       `json:"ID"`
	UserID    int       `json:"UserID"`
	PostID    int       `json:"PostID"`
	CreatedAt time.Time `json:"CreatedAt"`
}
