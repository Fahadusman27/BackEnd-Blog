package model

import "time"

type Komen struct {
	ID           int       `json:"id"`
	PostID      int       `json:"PostID"`
	UserID       int       `json:"UserID"`
	Comment_Text string    `json:"isi_komen"`
	Status       string    `json:"Status"`
	CreatedAt    time.Time `json:"CreatedAt"`
	UpdatedAt	time.Time	`json:"UpdatedAt"`
}
