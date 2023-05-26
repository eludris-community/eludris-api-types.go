package models

type ErrorMessage string

const (
	RateLimitedMessage ErrorMessage = "RATE_LIMITED"
	ValidationMessage  ErrorMessage = "VALIDATION"
	NotFoundMessage    ErrorMessage = "NOT_FOUND"
	ServerMessage      ErrorMessage = "SERVER"
)

// Represents an error returned by the Eludris API.
type BaseError struct {
	// The type of error that occurred.
	Type ErrorMessage `json:"type"`
	// The HTTP status code of the error.
	Status int `json:"status"`
	// A message which explains the error.
	Message ErrorMessage `json:"message"`
}

// The error raised when we are rate limited.
type RateLimitedError struct {
	BaseError
	// The number of milliseconds to wait before making new requests.
	RetryAfter int `json:"retry_after"`
}

// The error raised when the supplied request body is invalid.
type ValidationError struct {
	BaseError
	// The name of the value that raised the error.
	ValueName string `json:"value_name"`
	// The error that occurred.
	Error string `json:"error"`
	// Additional information about the error.
	Info string `json:"info"`
}

// The error raised when the requested resource could not be found.
type NotFoundError struct {
	BaseError
}

// The error raised when an internal server error occurs.
type ServerError struct {
	BaseError
	// Additional information about the error.
	Info string `json:"info"`
}
