package model

import "time"

type Komen struct {
	ID           int       `json:"id"`
	Post_ID      int       `json:"Post_id"`
	UserID       int       `json:"UserID"`
	Comment_Text string    `json:"isi_komen"`
	Status       string    `json:"Status"`
	CreatedAt    time.Time `json:"creadted_at"`
}
