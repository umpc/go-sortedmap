package desc

import "time"

func Time(i, j interface{}) bool {
	return i.(time.Time).After(j.(time.Time))
}