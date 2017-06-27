package sortedmap

import "sort"

func (sm *SortedMap) insertSort(key, val interface{}) []interface{} {
	smLen := len(sm.sorted)
	if smLen == 0 {
		return []interface{}{key}
	}
	return insertInterface(sm.sorted, key, sort.Search(smLen, func(i int) bool {
		return sm.lessFn(val, sm.idx[sm.sorted[i]])
	}))
}