package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wisnunu254/api-auth-golang/pkg/db"
	"github.com/wisnunu254/api-auth-golang/router"
)

func Server() {

	// connect to DB
	if err := db.Connects(); err != nil {
		log.Panicf("failed database setup. error: %v", err)
	}

	// Start the fiber server
	app := fiber.New()
	apiV1 := app.Group("/api/v1")
	db.Connects()
	// Use the AuthRouter function to define routes for the /api/v1 group
	router.AuthRouter(apiV1.(*fiber.Group))

	app.Listen(":3000")
}
