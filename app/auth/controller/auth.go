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

// AuthLogin handles the authentication process.
func AuthLogin(c *fiber.Ctx) error {
	// Parse request body
	login := new(model.AuthModelLogin)
	if err := c.BodyParser(login); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}
	// Retrieve user from the database
	// Retrieve user from the database
	user, err := getUserByEmail(login.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(http.StatusInternalServerError, response.MsgFailed, err.Error()))
	}

	// Check password
	if !checkPassword(user.Password, login.Password) {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(http.StatusBadRequest, response.MsgFailed, "Password does not match"))
	}

	// Generate JWT tokens
	jwtTokens, err := generateAccessToken(user.ID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	// Prepare response
	responseData := fiber.Map{
		"user":  dto.UserToDTO(user),
		"token": dto.TokenToDTO(jwtTokens),
	}

	return c.Status(http.StatusOK).JSON(response.ResponsSuccess(http.StatusOK, response.MsgSuccess, responseData))
}

func AuthRegister(c *fiber.Ctx) error {
	// Parse request body
	register := new(model.AuthModelRegister)
	if err := c.BodyParser(register); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	// Check if user already exists
	if checking, _ := getUserByEmail(register.Email); checking != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(http.StatusBadRequest, response.MsgFailed, "Email already exists"))
	}

	encryptedPassword, err := encryptPassword([]byte(register.Password))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	// Create user model
	user := &model.UserInsert{
		Email:    register.Email,
		Password: encryptedPassword,
	}

	// Insert user into the database
	userRepo := repo.UsersRepositorys(db.StartDB())
	if err := userRepo.InsertUsersRepository(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	// Respond with success
	return c.Status(http.StatusOK).JSON(response.ResponsSuccess(http.StatusOK, response.MsgSuccess, nil))
}

// getUserByEmail retrieves a user from the database by email.
func getUserByEmail(email string) (*model.User, error) {
	userRepo := repo.UsersRepositorys(db.StartDB())
	return userRepo.GetEmailUsersRepository(email)
}

// checkPassword compares the hashed password with the provided password.
func checkPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func encryptPassword(password []byte) (string, error) {
	// Generate hash from password
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

// generateAccessToken creates a JWT token for the given user ID.
func generateAccessToken(userID int64) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.ConfigApp().JwtSecretExpire)).Unix()

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(config.ConfigApp().JwtSecretKey))
}
