package adapter

// Error error struct
type Error struct {
	Message string
}

// NewError error message
func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}
