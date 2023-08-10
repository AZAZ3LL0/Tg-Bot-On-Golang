package errors

import "errors"

var (
	ErrBookNotAvailable = errors.New("book is not available for borrowing")
)
