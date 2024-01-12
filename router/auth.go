package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisnunu254/api-auth-golang/app/auth/controller"
)

func AuthRouter(apiV1 *fiber.Group) {
	route := apiV1.Group("/auth")
	route.Post("/login", controller.AuthLogin)
	route.Post("/register", controller.AuthRegister)
}
