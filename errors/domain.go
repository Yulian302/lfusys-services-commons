package errors

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidToken       = errors.New("invalid token")
	ErrInvalidTokenType   = errors.New("invalid token type")
	ErrTokenSignature     = errors.New("could not sign token")
	ErrTokenCreation      = errors.New("could not create token")

	ErrInternalServer = errors.New("internal server error")

	// uploads
	ErrFileSizeExceeded           = errors.New("file size exceeded")
	ErrFileSizeInvalid            = errors.New("invalid file size")
	ErrSessionConflict            = errors.New("session already exists")
	ErrGrpcFailed                 = errors.New("grpc call failed")
	ErrSessionNotFound            = errors.New("session not found")
	ErrFileNotFound               = errors.New("file not found")
	ErrSessionUpdateDetails       = errors.New("could not update session details")
	ErrUploadCompleteNotifyFailed = errors.New("upload complete notification failed")

	ErrServiceUnavailable = errors.New("service unavailable")
)
