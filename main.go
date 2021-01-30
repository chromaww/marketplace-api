package main

import (
	"marketplace-api/application/router"
	"marketplace-api/config"
	"marketplace-api/server"
	"os"
)

func main() {
	os.Setenv("APP_MODE", "development")

	cfg := config.Load(os.Getenv("APP_MODE"))

	os.Setenv("APP_VERSION", cfg.Version)

	app := router.Init(cfg.Version)

	server.Start(cfg, app)
}
