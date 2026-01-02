package errors

import (
	"log"

	"github.com/gin-gonic/gin"
)

type HTTPError struct {
	Code  int
	Error string
}

func NewHTTPError(message string, code int) *HTTPError {
	return &HTTPError{Error: message, Code: code}
}

func NewUnauthorizedError(msg string) HTTPError {
	return HTTPError{Error: msg, Code: 401}
}

func NewNotFoundError(msg string) HTTPError {
	return HTTPError{Error: msg, Code: 404}
}

func NewConflictError(msg string) HTTPError {
	return HTTPError{Error: msg, Code: 409}
}

func NewBadRequestError(msg string) HTTPError {
	return HTTPError{Error: msg, Code: 400}
}

func NewInternalServerError(msg string) HTTPError {
	return HTTPError{Error: msg, Code: 500}
}

func NewServiceUnavailableError(msg string) HTTPError {
	return HTTPError{Error: msg, Code: 503}
}

func NewForbiddenError(msg string) HTTPError {
	return HTTPError{
		Error: msg, Code: 403,
	}
}

// responses

func JSONErrorResponse(ctx *gin.Context, httpError HTTPError) {
	ctx.JSON(httpError.Code, gin.H{
		"error": httpError.Error,
	})
}

func InternalServerErrorResponse(ctx *gin.Context, msg string) {
	log.Printf("Internal server error: %v", msg)
	if gin.Mode() == gin.DebugMode {
		JSONErrorResponse(ctx, NewInternalServerError(msg))
	} else {
		JSONErrorResponse(ctx, NewInternalServerError("internal server error")) // generic in PROD for security
	}
}

func UnauthorizedResponse(ctx *gin.Context, msg string) {
	JSONErrorResponse(ctx, NewUnauthorizedError(msg))
}

func BadRequestResponse(ctx *gin.Context, msg string) {
	JSONErrorResponse(ctx, NewBadRequestError(msg))
}

func ConflictResponse(ctx *gin.Context, msg string) {
	JSONErrorResponse(ctx, NewConflictError(msg))
}

func ServiceUnavailableResponse(ctx *gin.Context, msg string) {
	JSONErrorResponse(ctx, NewServiceUnavailableError(msg))
}

func NotFoundResponse(ctx *gin.Context, msg string) {
	JSONErrorResponse(ctx, NewNotFoundError(msg))
}

func ForbiddenResponse(ctx *gin.Context, msg string) {
	JSONErrorResponse(ctx, NewForbiddenError(msg))
}
