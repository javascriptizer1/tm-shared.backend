package domainerror

// errorWithClass wraps an error with a specific class.
type errorWithClass struct {
	error
	class ErrorClass
	code  string
}

// ErrorClass returns the class of the error.
func (e *errorWithClass) ErrorClass() ErrorClass {
	return e.class
}

// ErrorCode returns the code of the error.
func (e *errorWithClass) ErrorCode() string {
	return e.code
}

// parse extracts the class, code, and message from an error.
func parse(err error) (class ErrorClass, code, message string) {
	if err == nil {
		panic("errs: nil error")
	}

	class = Internal
	if impl, ok := err.(interface{ ErrorClass() ErrorClass }); ok {
		class = impl.ErrorClass()
	}

	code = "Unknown"
	if impl, ok := err.(interface{ ErrorCode() string }); ok {
		code = impl.ErrorCode()
	}

	message = err.Error()

	return class, code, message
}
