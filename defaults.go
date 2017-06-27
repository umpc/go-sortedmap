package sortedmap

import "time"

func defaultSortLessFn(idx map[string]interface{}, sorted []string, i int, val interface{}) bool {
	return idx[sorted[i]].(time.Time).Before(val.(time.Time))
}

func setDefaults(lessFn SortLessFunc) SortLessFunc {
	if lessFn == nil {
		lessFn = defaultSortLessFn
	}
	return lessFn
}