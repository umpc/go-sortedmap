package sortedmap

// Get retrieves a value from the collection, using the given key.
func (sm *SortedMap) Get(key interface{}) (interface{}, bool) {
	val, ok := sm.idx[key]
	return val, ok
}

// BatchGet retrieves values with their read statuses from the collection, using the given keys.
func (sm *SortedMap) BatchGet(keys []interface{}) ([]interface{}, []bool) {
	vals := make([]interface{}, len(keys))
	results := make([]bool, len(keys))

	for i, key := range keys {
		vals[i], results[i] = sm.idx[key]
	}

	return vals, results
}

// GetMap returns a map containing keys mapped to values.
// The returned map is valid until the next modification to the SortedMap structure.
func (sm *SortedMap) GetMap() map[interface{}]interface{} {
	return sm.idx
}

// GetKeys returns a slice containing sorted keys.
// The returned slice is valid until the next modification to the SortedMap structure.
func (sm *SortedMap) GetKeys() []interface{} {
	return sm.sorted
}