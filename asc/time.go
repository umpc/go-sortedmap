package asc

import "time"

// Time is a less than comparison function for the time.Time type.
func Time(i, j interface{}) bool {
	return i.(time.Time).Before(j.(time.Time))
}
