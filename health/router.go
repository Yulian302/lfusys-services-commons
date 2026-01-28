package health

import (
	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(h *HealthHandler, r *gin.Engine) {
	health := r.Group("/health")

	health.GET("/live", h.Live)
	health.GET("/ready", h.Ready)
}
