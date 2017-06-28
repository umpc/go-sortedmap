package sortedmap

import "sort"

func (sm *SortedMap) insertSort(key, val interface{}) []interface{} {
	return insertInterface(sm.sorted, key, sort.Search(len(sm.sorted), func(i int) bool {
		return sm.lessFn(val, sm.idx[sm.sorted[i]])
	}))
}