package main

import (
	"BLOG/domain/repository"
	"BLOG/domain/route"
	"BLOG/domain/service"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := &sql.DB{}

	Authrepo := repository.NewAuthRepository(db)
	CommentRepo := repository.NewCommentRepository(db)
	LikeRepo := repository.NewLikeRepository(db)
	ShareRepo := repository.NewShareRepository(db)
	SearchRepo := repository.NewSearchRepository(db)
	
	auth := service.NewAuthService(Authrepo, "SECRET_KEY_ANDA")
	comment := service.NewCommentService(CommentRepo)
	Like := service.NewLikeService(LikeRepo)
	Share := service.NewShareService(ShareRepo)
	search := service.NewSearchService(SearchRepo)

	route.AutRoute(app, auth)
	route.PostRoute(app, comment, Like, Share)
	route.SearchRoute(app, search)

    app.Listen(":8080")
}