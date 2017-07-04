package desc

import "time"

// Time is a greater than comparison function for the time.Time type.
func Time(i, j interface{}) bool {
	return i.(time.Time).After(j.(time.Time))
}
