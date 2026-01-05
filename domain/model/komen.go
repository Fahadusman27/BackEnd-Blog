package model

import "time"


type Komen struct {
	ID	int	`json:"id"`
	Post_ID	[]Post	`json:"Post_id"`
	Name	string	`json:"Name"`
	Comment_Text	string	`json:"Komen"`
	Status	string	`json:"Status"`
	CreatedAt	time.Time	`json:"CreadtedAt"`
}