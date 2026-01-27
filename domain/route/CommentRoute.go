package route

import (
	."BLOG/domain/service"

	"github.com/gofiber/fiber/v2"
)

func CommentRoute(app *fiber.App, k *KomenService) {
	comment := app.Group("/")

	comment.Post("/Comment", k.CommentService)
}