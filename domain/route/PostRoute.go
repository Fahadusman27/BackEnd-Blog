package route

import (
	."BLOG/domain/service"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App, k *KomenService, l *LikeService, s *ShareService) {
	Post := app.Group("/Post")

	Post.Post("/Comment", k.CommentService)
	Post.Post("/Like/PostID", l.LikeAction)
	Post.Post("/Share/PostID", s.ShareAction)
}