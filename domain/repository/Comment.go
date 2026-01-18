package repository

import (
	"BLOG/domain/config"
	. "BLOG/domain/model"
)

func CreateCommet(ID int) (*Komen, error) {
	komen := new(Komen)

	query := `INSERT INTO Komen (post_id, isi_komen, UserID, Status) 
	VALUES ($1, $2, $3, $4)`

	_, err := config.DB.Exec(query, &komen.Post_ID, &komen.Comment_Text, &komen.UserID, &komen.Status)

	if err != nil {
		return nil, err
	}

	return komen, nil
}