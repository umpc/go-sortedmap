package sortedmap

import (
	"errors"
	"sort"
)

func (sm *SortedMap) delete(key interface{}) bool {
	if val, ok := sm.idx[key]; ok {

		smLen := len(sm.sorted)
		i := sort.Search(smLen, func(i int) bool {
			return sm.lessFn(val, sm.idx[sm.sorted[i]])
		})

		if i == smLen {
			i--
		} else if i < smLen-1 {
			i++
		}
		for sm.sorted[i] != key {
			i--
		}

		delete(sm.idx, key)
		sm.sorted = deleteInterface(sm.sorted, i)

		return true
	}
	return false
}

func (sm *SortedMap) boundedDelete(lowerBound, upperBound interface{}) error {
	iterBounds := sm.boundsIdxSearch(lowerBound, upperBound)
	if iterBounds == nil {
		return errors.New(noValuesErr)
	}
	for i, deleted := iterBounds[0], 0; i <= iterBounds[1]-deleted; i++ {
		delete(sm.idx, sm.sorted[i])
		sm.sorted = deleteInterface(sm.sorted, i)
		deleted++
	}
	return nil
}

// Delete removes a value from the collection, using the given key.
// Because the index position of each sorted key changes on each insert and a simpler structure was ideal, deletes can have a worse-case complexity of O(n), meaning the goroutine must loop through the sorted slice to find and delete the given key.
func (sm *SortedMap) Delete(key interface{}) bool {
	return sm.delete(key)
}

// BatchDelete removes values from the collection, using the given keys, returning a slice of the results.
func (sm *SortedMap) BatchDelete(keys []interface{}) []bool {
	results := make([]bool, len(keys))
	for i, key := range keys {
		results[i] = sm.delete(key)
	}
	return results
}

// BoundedDelete removes values that are between the given values from the collection.
// BoundedDelete returns true if the operation was successful, or false otherwise.
func (sm *SortedMap) BoundedDelete(lowerBound, upperBound interface{}) error {
	return sm.boundedDelete(lowerBound, upperBound)
}
