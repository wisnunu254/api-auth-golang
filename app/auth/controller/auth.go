package controller

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/wisnunu254/api-auth-golang/app/auth/dto"
	"github.com/wisnunu254/api-auth-golang/app/auth/model"
	"github.com/wisnunu254/api-auth-golang/config"
	"github.com/wisnunu254/api-auth-golang/pkg/db"
	repo "github.com/wisnunu254/api-auth-golang/repository"
	"github.com/wisnunu254/api-auth-golang/util/response"
	"golang.org/x/crypto/bcrypt"
)

func AuthLogin(c *fiber.Ctx) error {
	// Parse request body
	login := new(model.AuthModelLogin)
	if err := c.BodyParser(login); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error(),
		))
	}

	// Retrieve user from the database
	userRepo := repo.UsersRepositorys(db.StartDB())
	user, err := userRepo.GetEmailUsersRepository(login.Email)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error(),
		))
	}

	// Check password
	if !checkingPassword([]byte(user.Password), []byte(login.Password)) {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			"password not match",
		))
	}

	// Generate JWT tokens
	jwtTokens, err := AccessToken(user.ID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error(),
		))
	}

	// Prepare response
	responseData := fiber.Map{
		"data": fiber.Map{
			"user":  dto.UserToDTO(user),
			"token": dto.TokenToDTO(jwtTokens),
		},
	}

	return c.Status(http.StatusOK).JSON(response.ResponsSuccess(
		http.StatusOK,
		response.MsgSuccess,
		responseData,
	))
}

func AuthRegister() {

}

func AuthRefreshTokens() {

}
func AccessToken(userID int64) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.ConfigApp().JwtSecretExpire)).Unix()

	// Generate encoded token and send it as response.
	tokens, err := token.SignedString([]byte(config.ConfigApp().JwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokens, nil
}
func checkingPassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		// Check if the error is due to an incorrect password
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false
		}
		// Handle other potential errors
		return false
	}

	return true
}
