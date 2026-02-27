package middleware

import (
	"strings"

	"github.com/Yulian302/lfusys-services-commons/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApplyCors(r *gin.Engine, cfg config.CorsConfig) {
	corsCfg := cors.Config{
		AllowOrigins:     strings.Split(cfg.Origins, ","),
		AllowMethods:     strings.Split(cfg.Methods, ","),
		AllowHeaders:     strings.Split(cfg.Headers, ","),
		AllowCredentials: cfg.Credentials,
	}

	r.Use(cors.New(
		corsCfg,
	))
}
