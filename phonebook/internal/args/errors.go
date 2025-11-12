package args

import "errors"

var (
	ErrInvalidUsage     = errors.New("invalid usage")
	ErrInvalidArguments = errors.New("invalid arguments")
	ErrInvalidCommand   = errors.New("invalid command")
)
