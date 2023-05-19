package implementations

import "github.com/pkg/errors"

var (
	ErrLimitIsZero  = errors.New("the limit can't be 0")
	ErrOrderByIsNil = errors.New("the parameter OrderBy cannot be nil")
)
