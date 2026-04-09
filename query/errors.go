package query

import "errors"

var (
	ErrColumnNotFound = errors.New("column not found")
	ErrRowNotFound    = errors.New("row not found")
)
