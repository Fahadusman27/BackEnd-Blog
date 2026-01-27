package service

import (
	. "BLOG/domain/model"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type KomenService struct {
	repo CommentRepo
}

type CommentRepo interface {
	CreateComment(UserID int, PostID int, content string) (*Komen, error)
}

func NewCommentService(repo CommentRepo) *KomenService {
	return &KomenService{repo: repo}
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