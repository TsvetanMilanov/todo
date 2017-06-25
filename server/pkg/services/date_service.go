package services

import (
	"time"
)

// DateService date/time realted methods.
type DateService struct {
}

// Now returns the current time.
func (date *DateService) Now() time.Time {
	return time.Now()
}
