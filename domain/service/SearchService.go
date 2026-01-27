package service

import (
	. "BLOG/domain/model"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type SearchService struct {
	repo SearchRepo
}

type SearchRepo interface {
	SeacrhByID(ID int) (*Users, error)
	SearchByUsername(username string) (*Users, error)
	SearchByKategori(Name string) (*Kategori, error)
	SearchByTitle(Name string) (*Title, error)
}

func NewSearchService(repo SearchRepo) *SearchService {
	return &SearchService{repo: repo}
}

func (s *SearchService) Search(keyword string) (map[string]interface{}, error) {
	if keyword == "" {
		return nil, errors.New("Kata Kunci tidak ada")
	}

	var userFound interface{}
	var kategoriFound interface{}
	var titleFound interface{}

	u, err := s.repo.SearchByUsername(keyword)
	if err == nil {
		userFound = u
	}

	k, err := s.repo.SearchByKategori(keyword)
	if err == nil {
		kategoriFound = k
	}

	t, err := s.repo.SearchByTitle(keyword)
	if err == nil {
		titleFound = t
	}

	return map[string]interface{}{
		"users":      userFound,
		"categories": kategoriFound,
		"articles":   titleFound,
		"query":      keyword,
	}, nil
}

func (s *SearchService) SearchServices(c *fiber.Ctx) error {
	keyword := c.Query("q")

	result, err := s.Search(keyword)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Hasil Pencarian",
		"data":    result,
	})
}
