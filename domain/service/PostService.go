package service

import (
	. "BLOG/domain/model"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Komentar

type KomenService struct {
	repo CommentRepo
}

type CommentRepo interface {
	CreateComment(UserID int, PostID int, content string) (*Komen, error)
}

func NewCommentService(repo CommentRepo) *KomenService {
	return &KomenService{}
}

func IsDirty(content string) bool {
	badWords := []string{"anjing", "babi", "bangke", "goblok", "tolol", "bego",
	"sinting", "gila", "kurang ajar", "sialan", "brengsek",
	"kampret", "bedebah", "kontol", "memek", "ngentot",
	"tai", "bangsat", "perek", "lonte", "monyet", "kera",
	"asu", "jancok", "bajingan", "cok", "ajig", "kehed", "bacot"}

	for _,word := range badWords {
		if strings.Contains(strings.ToLower(content), word) {
			return true
		}
	}

	return false
}

func (c *KomenService) Comment(UserID int, PostID int, content string) (*Komen, error) {
	if strings.TrimSpace(content) == "" {
		return nil, errors.New("komentar tidak boleh kosong")
	}
	
	if IsDirty(content) {
		return nil, errors.New("Komentar Mengandung Kata Tidak Pantas")
	}

	res, err := c.repo.CreateComment(UserID, PostID, content)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (k *KomenService) CommentService(c *fiber.Ctx) error {
	var comment Komen

	if err := c.BodyParser(&comment); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error" : "Gagal mengirim komentar, silakan coba lagi.",
		})
	}

	if strings.TrimSpace(comment.Comment_Text) == "" {
		return c.Status(400).JSON(fiber.Map{
			"error" : "tidak boleh kosong!",
		})
	}

	if IsDirty(comment.Comment_Text) {
		return c.Status(400).JSON(fiber.Map{
			"pesan" : "tidak boleh menggunakan kalimat jorok atau kasar",
		})
	}
	
	return c.JSON(fiber.Map{
		"pesan" : "Komentar berhasil terkirim",
	})
}


// Like

type LikeService struct {
	repo LikeRepo
}

type LikeRepo interface {
	CreateLike(UserID, PostID int) (string, error)
	CreateStatistikLike(PostID int) (int, error)
}

func NewLikeService(repo LikeRepo) *LikeService {
	return &LikeService{repo : repo}
}

func (l *LikeService) LikeAction(c *fiber.Ctx) error {
	PostID, _ := c.ParamsInt("PostID")

	UserID, _ := c.Locals("UserID").(int)

	status, err := l.repo.CreateLike(UserID, PostID)

	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"error" : "Gagal Like postingan",
		})
	}

	total, _ := l.repo.CreateStatistikLike(PostID)

	return c.JSON(fiber.Map{
		"pesan" : status,
		"total like" : total,
	})
}

// Share

type ShareService struct {
	repo ShareRepo
}

type ShareRepo interface {
	PostShare(UserID, PostID int, Platform string) error
	CreateStatistikShare(PostID int) (int, error)
}

func NewShareService(repo ShareRepo) *ShareService {
	return &ShareService{repo : repo}
}

func (s *ShareService) ShareAction(c *fiber.Ctx) error {
	PostID, _ := c.ParamsInt("PostID")

	Platform := c.Query("p", "general")

	UserID := 1

	if err := s.repo.PostShare(UserID, PostID, Platform); err != nil{
		return c.Status(500).JSON(fiber.Map{
			"error" : "Gagal Share",
		})
	}

	total, _ := s.repo.CreateStatistikShare(PostID)

	return c.JSON(fiber.Map{
		"pesan" : "Berhasil share ke " + Platform,
		"total share" : total,
	})
}