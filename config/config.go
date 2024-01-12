package config

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func ConfigsEnv(envFiles string) {
	err := godotenv.Load(envFiles)
	if err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}
	ConfigsApp()
	ConfigsDB()
}

func ConfigsFiber() fiber.Config {
	return fiber.Config{
		// Prefork:       true,
		// CaseSensitive: true,
		// StrictRouting: true,
		// ServerHeader:  "Fiber",
		// AppName:       "Test App v1.0.1",
		ReadTimeout: time.Second * time.Duration(ConfigApp().ReadTimeOut),
	}
}
