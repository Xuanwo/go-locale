package locale

// Error is the error returned by locale.
type Error struct {
	Op  string
	Err error
}

func (e *Error) Error() string {
	return e.Op + ": " + e.Err.Error()
}

// Unwrap implements xerrors.Wrapper
func (e *Error) Unwrap() error {
	return e.Err
}
