package route

import (
	."BLOG/domain/service"

	"github.com/gofiber/fiber/v2"
)

func SearchRoute(app *fiber.App, s *SearchService) {
	search := app.Group("/Search")

	search.Get("/", s.SearchServices)
}