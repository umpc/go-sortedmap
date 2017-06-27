package sortedmap

import "sort"

func (sm *SortedMap) insertSort(key string, val interface{}) []string {
	smLen := len(sm.sorted)
	if smLen == 0 {
		return []string{key}
	}
	i := sort.Search(smLen, func(i int) bool {
		return sm.lessFn(val, sm.idx[sm.sorted[i]])
	})
	return insertString(sm.sorted, i, key)
}