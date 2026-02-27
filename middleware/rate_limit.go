package middleware

import (
	"time"

	logger "github.com/Yulian302/lfusys-services-commons/logging"
	"github.com/Yulian302/lfusys-services-commons/ratelimit"
	"github.com/Yulian302/lfusys-services-gateway/middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func ApplyRateLimiting(r *gin.Engine, client *redis.Client, l logger.Logger) {
	rateLimiter := ratelimit.NewRedisRateLimiter(client)
	r.Use(middleware.RateLimiterMiddleware(rateLimiter, 100, time.Minute, l))
}
