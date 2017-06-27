package sortedmap

func (sm *SortedMap) delete(key interface{}) bool {
	if _, ok := sm.idx[key]; ok {
		delete(sm.idx, key)

		smLen := len(sm.sorted)
		deleted := 0

		for i := 0; i < smLen - deleted; i++ {
			if sm.sorted[i] == key {
				sm.sorted = deleteInterface(sm.sorted, i)
				deleted++
			}
		}

		return true
	}
	return false
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