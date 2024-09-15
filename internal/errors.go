package internal

import "errors"

var (
	ErrNoRecord       = errors.New("ErrNoRecord: the requested record was not found")
	ErrDuplicateEmail = errors.New("ErrDuplicateEmail: duplicate email address")
)
