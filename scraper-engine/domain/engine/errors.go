package engine

import "errors"

var (
	ErrInvalidTaskType       = errors.New("invalid task type! Unknown task type")
	ErrEmptySearchKeyword    = errors.New("invalid seach keywords! Empty search")
	ErrEmptyTaskLocation     = errors.New("invalid location! Empty task location")
	ErrEmptyTaskLocationId   = errors.New("invalid location id! Empty task location id")
	ErrInvalidTaskLocationId = errors.New("invalid location id! Location Id must be numeric")
	ErrInvalidDistanceType   = errors.New("invalid distance type! Not an integer")
	ErrInvalidDelay          = errors.New("invalid delay! less than 30 minutes")
	ErrInvalidDistance       = errors.New("invalid distance has been passed! distance should be in between 25 and 100")
	ErrNotFound              = errors.New("unable to find task by the Id")
)
