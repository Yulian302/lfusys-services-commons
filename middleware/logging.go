package middleware

import (
	"log/slog"

	"github.com/Yulian302/lfusys-services-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func ApplyLogging(r *gin.Engine, logger *slog.Logger) {
	r.Use(middleware.Logger(logger))
}
