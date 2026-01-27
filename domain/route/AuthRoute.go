package route

import (
	. "BLOG/domain/service"

	"github.com/gofiber/fiber/v2"
)

func AutRoute(app *fiber.App, a *AuthService) {
	auth := app.Group("/auth")

	auth.Post("/login", a.LoginService)
	auth.Get("/register", a.RegisterService)
	auth.Get("/profile", a.GetProfileService)
	auth.Post("/update/profile", a.UpdateProfileService)
}