package server

import (
	"marketplace-api/config"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config, r *gin.Engine) {
	// Default port set to be 8080
	var port string = "8080"

	if cfg.Port != 0 {
		port = strconv.Itoa(cfg.Port)
	}

	r.Run(":" + port)
}
