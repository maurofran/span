package span

// Error is the custom time for the span package
type Error string

// Error implements the built-in error interface.
func (e Error) Error() string {
	return string(e)
}

const (
	// ErrInvalidComparator is returned if no comparator is provided.
	ErrInvalidComparator = Error("invalid comparator")
	// ErrInvalidSpan is returned when a caller tries to create an invalid span.
	ErrInvalidSpan = Error("invalid span")
)
