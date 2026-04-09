package row

import "errors"

var (
	// ErrRowNotFound 行不存在
	ErrRowNotFound = errors.New("row not found")
	// ErrIndexOutOfRange 索引越界
	ErrIndexOutOfRange = errors.New("index out of range")
)
