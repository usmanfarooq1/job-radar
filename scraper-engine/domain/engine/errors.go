package engine

import "errors"

var (
	ErrInvalidTaskType     = errors.New("invalid task type! Unknown task type")
	ErrInvalidDistanceType = errors.New("invalid distance type! Not an integer")
	ErrInvalidDelay        = errors.New("invalid delay! less than 30 minutes")
	ErrInvalidDistance     = errors.New("invalid distance has been passed! distance should be in between 25 and 100")
	ErrNotFound            = errors.New("unable to find task by the Id")
)
