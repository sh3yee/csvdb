package column

import "errors"

var (
	// ErrColumnNotFound 列不存在
	ErrColumnNotFound = errors.New("column not found")
	// ErrIndexOutOfRange 索引越界
	ErrIndexOutOfRange = errors.New("index out of range")
)
