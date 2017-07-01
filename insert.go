package sortedmap

func (sm *SortedMap) insert(key, val interface{}) bool {
	if _, ok := sm.idx[key]; !ok {
		sm.idx[key] = val
		sm.sorted = sm.insertSort(key, val)
		return true
	}
	return false
}

// Insert uses the provided 'less than' function to insert sort and add the value to the collection and returns a value containing the record's insert status.
// If the key already exists, the value will not be inserted. Use Replace for the alternative functionality.
func (sm *SortedMap) Insert(key, val interface{}) bool {
	return sm.insert(key, val)
}

// BatchInsert adds all given records to the collection and returns a slice containing each record's insert status.
// If a key already exists, the value will not be inserted. Use BatchReplace for the alternative functionality.
func (sm *SortedMap) BatchInsert(recs []*Record) []bool {
	results := make([]bool, len(recs))
	for i, rec := range recs {
		results[i] = sm.insert(rec.Key, rec.Val)
	}
	return results
}

// BatchInsertMap adds all map keys and values to the collection and returns a slice containing each record's insert status.
// If a key already exists, the value will not be inserted. Use BatchReplaceMap for the alternative functionality.	
func (sm *SortedMap) BatchInsertMap(m map[interface{}]interface{}) []bool {
	results := make([]bool, len(m))
	i := 0
	for key, val := range m {
		results[i] = sm.insert(key, val)
		i++
	}
	return results
}