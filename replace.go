package sortedmap

func (sm *SortedMap) replace(key, val interface{}) {
	if _, ok := sm.idx[key]; ok {
	    sm.delete(key)
    }
	sm.idx[key] = val
	sm.sorted = sm.insertSort(key, val)
}

// Replace uses the provided 'less than' function to insert sort.
// Even if the key already exists, the value will be inserted. Use Insert for the alternative functionality.
func (sm *SortedMap) Replace(key, val interface{}) {
	sm.replace(key, val)
}

// BatchReplace adds all given records to the collection.
// Even if a key already exists, the value will be inserted. Use BatchInsert for the alternative functionality.	
func (sm *SortedMap) BatchReplace(recs []*Record) {
	for _, rec := range recs {
		sm.replace(rec.Key, rec.Val)
	}
}