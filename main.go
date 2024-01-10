package main

import (
	"github.com/wisnunu254/api-auth-golang/config"
	"github.com/wisnunu254/api-auth-golang/server"
)

// @title Fiber Go API
// @version 1.0
// @description Fiber go web framework based REST API boilerplate
// @contact.name H.R. Shadhin
// @contact.email dev@hrshadhin.me
// @termsOfService
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:5000
// @BasePath /api
func main() {

	config.ConfigsEnv(".env")
	server.Server()
}
