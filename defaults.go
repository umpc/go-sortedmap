package sortedmap

import "time"

func defaultSortLessFn(insertVal, idxVal interface{}) bool {
	return insertVal.(time.Time).Before(idxVal.(time.Time))
}

func setDefaults(lessFn SortLessFunc) SortLessFunc {
	if lessFn == nil {
		lessFn = defaultSortLessFn
	}
	return lessFn
}