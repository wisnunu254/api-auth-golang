package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wisnunu254/api-auth-golang/config"
	"github.com/wisnunu254/api-auth-golang/pkg/db"
	"github.com/wisnunu254/api-auth-golang/router"
)

func Server() {

	// connect to DB
	if err := db.Connects(); err != nil {
		log.Panicf("failed database setup. error: %v", err)
	}
	db.Connects()
	// Start the fiber server
	fiberConfig := config.ConfigsFiber()
	app := fiber.New(fiberConfig)
	apiV1 := app.Group("/api/v1")
	// Use the AuthRouter function to define routes for the /api/v1 group
	router.AuthRouter(apiV1.(*fiber.Group))
	defer db.StartDB().Close()
	app.Listen(":3000")
}
