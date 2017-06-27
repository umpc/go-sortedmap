package sortedmap

func (sm *SortedMap) replace(key string, val interface{}) {
	if _, ok := sm.idx[key]; ok {
	    sm.delete(key)
    }
	sm.idx[key] = val
	sm.sorted = sm.insertSort(key, val)
}

// Replace uses the provided 'less than' function to insert sort.
// Even if the key already exists, the value will be inserted. Use Insert for the alternative functionality.
func (sm *SortedMap) Replace(key string, val interface{}) {
	sm.replace(key, val)
}

// BatchReplace adds all given records to the collection.
// Even if a key already exists, the value will be inserted. Use BatchInsert for the alternative functionality.	
func (sm *SortedMap) BatchReplace(recs ...*Record) {
	for _, rec := range recs {
		sm.replace(rec.Key, rec.Val)
	}
}

// ChReplace reads records from a channel and adds all given records to the collection.
// Even if a key already exists, the value will be inserted. If a clearly efficient alternative to this function is proposed, it will likely be accepted and merged.
func (sm *SortedMap) ChReplace(ch <-chan *Record) {
	for rec := range ch {
		sm.replace(rec.Key, rec.Val)
	}
}