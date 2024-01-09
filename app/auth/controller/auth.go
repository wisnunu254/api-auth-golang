package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wisnunu254/api-auth-golang/app/auth/model"
	"github.com/wisnunu254/api-auth-golang/util/response"
)

func AuthLogin(c *fiber.Ctx) error {
	login := &model.AuthModelLogin{}

	if err := c.BodyParser(login); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error(),
		))
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
	})
}

func AuthRegister() {

}

func AuthRefreshTokens() {

}
