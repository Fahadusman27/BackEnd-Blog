package repository

import (
	"BLOG/domain/config"
	. "BLOG/domain/model"
	"database/sql"
)

// Komentar

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (k *CommentRepository) CreateComment(UserID int, PostID int, content string) (*Komen, error) {
	query := `INSERT INTO komen (PostID, isi_komen, UserID, status)
	VALUES ($1, $2, $3, $4)
	RETURNING id, PostID, isi_komen, UserID, status, CreatedAt, UpdateAt`

	komen := new(Komen)

	err := config.DB.QueryRow(query, PostID, content, UserID, "aktif").
		Scan(&komen.ID, &komen.PostID, &komen.Comment_Text, &komen.UserID, komen.Status, &komen.CreatedAt)

	if err != nil {
		return nil, err
	}

	return komen, nil
}

// Like

type LikeRepository struct {
	db *sql.DB
}

func NewLikeRepository(db *sql.DB) *LikeRepository {
	return &LikeRepository{db: db}
}

func (l *LikeRepository) CreateLike(UserID, PostID int) (string, error) {

	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM likes WHERE UserID=$1 AND PostID=$2)`
	l.db.QueryRow(checkQuery, UserID, PostID).Scan(&exists)

	if exists {
		query := `DELETE from likes WHERE UserID=$1 AND PostID=$2`
		_, err := l.db.Exec(query, UserID, PostID)
		return "Unlike", err
	} else {
		query := `INSERT INTO likes (UserID, PostID) VALUES ($1, $2)`
		_, err := l.db.Exec(query, UserID, PostID)
		return "Liked", err
	}
}

func (s *LikeRepository) CreateStatistikLike(PostID int) (int, error) {
	query := `SELECT COUNT(*) FROM likes WHERE PostID = $1`

	var count int

	err := s.db.QueryRow(query, PostID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Share

type ShareRepository struct {
	db *sql.DB
}

func NewShareRepository(db *sql.DB) *ShareRepository {
	return &ShareRepository{db: db}
}

func (s *ShareRepository) PostShare(UserId, PostID int, Platform string) error {
	query := `INSERT INTO (PostID, UserID, Platform) VALUES ($1, $2, $3)`

	_, err := s.db.Exec(query, UserId, Platform)

	return err
}

func (s *ShareRepository) CreateStatistikShare(PostID int) (int, error) {
	query := `SELECT COUNT(*) FROM shares WHERE PostID = $1`

	var count int

	err := s.db.QueryRow(query, PostID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
