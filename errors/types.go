package errors

type HTTPError struct {
	Error string `json:"error" example:"error message"`
	Code  string `json:"code,omitempty" example:"INVALID_INPUT"`
}

func NewHTTPError(message string) HTTPError {
	return HTTPError{Error: message}
}
