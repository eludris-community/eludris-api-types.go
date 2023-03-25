package models

type ErrorMessage string

const (
	RateLimitedMessage         ErrorMessage = "You have been rate limited"
	FileSizeRateLimitedMessage ErrorMessage = "You have surpassed your file size limit"
	ValidationMessage          ErrorMessage = "Invalid request body"
	NotFoundMessage            ErrorMessage = "The requested resource cannot be found"
	ServerMessage              ErrorMessage = "The server encountered an error while performing the requested action"
)

// Represents an error returned by the Eludris API.
type BaseError struct {
	// The HTTP status code of the error.
	Status int `json:"status"`
	// A message which explains the error.
	Message ErrorMessage `json:"message"`
}

// The error raised when we are rate limited.
type RateLimitedError struct {
	BaseError
	Data struct {
		// The number of milliseconds to wait before making new requests.
		RetryAfter int `json:"retry_after"`
	} `json:"data"`
}

// The error raised when the maximum amount of bytes has been surpassed and we are rate limited.
type FileSizeRateLimitedError struct {
	BaseError
	Data struct {
		// The number of milliseconds to wait before making new requests.
		RetryAfter int `json:"retry_after"`
		// The number of bytes available until the rate limit resets.
		BytesLeft int `json:"bytes_left"`
	} `json:"data"`
}

// The error raised when the supplied request body is invalid.
type ValidationError struct {
	BaseError
	Data struct {
		// The name of the field that raised the error.
		FieldName string `json:"field_name"`
		// The error that occurred.
		Error string `json:"error"`
	} `json:"data"`
}

// The error raised when the requested resource could not be found.
type NotFoundError struct {
	BaseError
}

// The error raised when an internal server error occurs.
type ServerError struct {
	BaseError
	// The name of the field that raised the error.
	Error string `json:"error"`
}
