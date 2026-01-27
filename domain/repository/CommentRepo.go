package repository

import (
	"BLOG/domain/config"
	. "BLOG/domain/model"
	"database/sql"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (k *CommentRepository) CreateComment(UserID int, PostID int, content string) (*Komen, error) {
	query := `	INSERT INTO komen (PostID, isi_komen, UserID, status)
	VALUES ($1, $2, $3, $4)
	RETURNING id, PostID, isi_komen, UserID, status, CreatedAt, UpdateAt`

	komen := new(Komen)

	err := config.DB.QueryRow(query, PostID, content, UserID, "aktif",).
		Scan(&komen.ID, &komen.PostID, &komen.Comment_Text, &komen.UserID, komen.Status, &komen.CreatedAt)

	if err != nil {
		return nil, err
	}

	return komen, nil
}