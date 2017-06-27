package sortedmap

// Has checks if the key exists in the collection.
func (sm *SortedMap) Has(key string) bool {
	_, ok := sm.idx[key]
	return ok
}

// BatchHas checks if the keys exist in the collection and returns a slice containing the results.
func (sm *SortedMap) BatchHas(keys ...string) []bool {
	results := make([]bool, len(keys))
	for i, key := range keys {
		_, results[i] = sm.idx[key]
	}
	return results
}