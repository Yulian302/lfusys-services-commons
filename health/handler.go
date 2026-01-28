package health

import (
	"fmt"

	"github.com/Yulian302/lfusys-services-commons/errors"
	"github.com/Yulian302/lfusys-services-commons/responses"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	// critical services
	checks []ReadinessCheck
}

func NewHealthHandler(checks ...ReadinessCheck) *HealthHandler {
	return &HealthHandler{checks: checks}
}

func (h *HealthHandler) Live(c *gin.Context) {
	responses.JSONSuccess(c, "ok")
}

func (h *HealthHandler) Ready(c *gin.Context) {
	for _, check := range h.checks {
		if !check.IsReady() {
			errors.ServiceUnavailableResponse(c, fmt.Sprintf("service %s is not ready", check.Name()))
			return
		}
	}

	responses.JSONSuccess(c, "ready")
}
