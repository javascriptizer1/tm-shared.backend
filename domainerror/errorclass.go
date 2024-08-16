package domainerror

import "errors"

type ErrorClass string

const (
	// Unclassified class is for errors that can't be classified.
	Unclassified ErrorClass = "Unclassified"
	// Internal class indicates internal server errors.
	Internal ErrorClass = "Internal"
	// Invalid class indicates that the request is invalid.
	Invalid ErrorClass = "Invalid"
	// NotFound class indicates that the requested entity is not found.
	NotFound ErrorClass = "NotFound"
	// PermissionDenied class indicates permission issues.
	PermissionDenied ErrorClass = "PermissionDenied"
	// Unauthenticated class indicates authentication failures.
	Unauthenticated ErrorClass = "Unauthenticated"
	// Conflict class indicates resource conflict.
	Conflict ErrorClass = "Conflict"
	// Unavailable class indicates that the service is unavailable.
	Unavailable ErrorClass = "Unavailable"
	// Unimplemented class indicates that the operation is not implemented.
	Unimplemented ErrorClass = "Unimplemented"
)

// New creates a new error with the specified class, code, and message.
func (c ErrorClass) New(code, message string) error {
	return &errorWithClass{
		error: errors.New(message),
		class: c,
		code:  code,
	}
}

// As wraps an existing error with a class and code.
func (c ErrorClass) As(code string, err error) error {
	if err == nil {
		panic("errs: nil error")
	}

	return &errorWithClass{
		error: err,
		class: c,
		code:  code,
	}
}
