package asc

import "time"

func Time(i, j interface{}) bool {
	return i.(time.Time).Before(j.(time.Time))
}
