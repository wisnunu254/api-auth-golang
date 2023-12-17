package config

import (
	"os"
	"strconv"
	"time"
)

type App struct {
	Host        string
	Port        int
	Debug       bool
	ReadTimeOut time.Duration

	// JWT Config APP
	JwtSecretKey    string
	JwtSecretExpire int
}

var app = &App{}

func ConfigApp() *App {
	return app
}

func ConfigsApp() {
	app.Host = os.Getenv("APP_HOST")
	app.Port, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	app.Debug = os.Getenv("APP_DEBUG") == "true"
	app.ReadTimeOut = 5 * time.Second

	app.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	app.JwtSecretExpire, _ = strconv.Atoi(os.Getenv("JWT_SECRET_EXPIRE"))
}
