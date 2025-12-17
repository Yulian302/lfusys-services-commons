package errors

import (
	"log"

	"github.com/gin-gonic/gin"
)

func JSONError(ctx *gin.Context, statusCode int, msg string) {
	ctx.JSON(statusCode, NewHTTPError(msg))
}

func InternalServerError(ctx *gin.Context, msg string) {
	log.Printf("Internal server error: %v", msg)
	if gin.Mode() == gin.DebugMode {
		JSONError(ctx, 500, msg)
	} else {
		JSONError(ctx, 500, "internal server error") // generic in PROD for security
	}
}

func BadRequestError(ctx *gin.Context, msg string) {
	JSONError(ctx, 400, msg)
}

func Unauthorized(ctx *gin.Context, msg string) {
	JSONError(ctx, 401, msg)
}
