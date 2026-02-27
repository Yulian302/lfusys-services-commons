package health

import (
	"fmt"

	"github.com/Yulian302/lfusys-services-commons/errors"
	logger "github.com/Yulian302/lfusys-services-commons/logging"
	"github.com/Yulian302/lfusys-services-commons/responses"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	// critical services
	checks []ReadinessCheck

	logger logger.Logger
}

func NewHealthHandler(l logger.Logger, checks ...ReadinessCheck) *HealthHandler {
	return &HealthHandler{checks: checks, logger: l}
}

func (h *HealthHandler) Live(c *gin.Context) {
	responses.JSONSuccess(c, "ok")
}

func (h *HealthHandler) Ready(c *gin.Context) {
	ctx := c.Request.Context()

	for _, check := range h.checks {
		if err := check.IsReady(ctx); err != nil {
			h.logger.Error(fmt.Sprintf("service %s is not ready", check.Name()), "reason", err.Error())
			errors.ServiceUnavailableResponse(c, fmt.Sprintf("service %s is not ready", check.Name()))
			return
		}
	}

	responses.JSONSuccess(c, "ready")
}
