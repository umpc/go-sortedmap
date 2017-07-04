package sortedmap

import "fmt"

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
func (sm *SortedMap) BatchInsert(recs []Record) []bool {
	results := make([]bool, len(recs))
	for i, rec := range recs {
		results[i] = sm.insert(rec.Key, rec.Val)
	}
	return results
}

func (sm *SortedMap) batchInsertMapInterfaceKeys(m map[interface{}]interface{}) error {
	for key, val := range m {
		if !sm.insert(key, val) {
			return fmt.Errorf("Key already exists: %+v", key)
		}
	}
	return nil
}

func (sm *SortedMap) batchInsertMapStringKeys(m map[string]interface{}) error {
	for key, val := range m {
		if !sm.insert(key, val) {
			return fmt.Errorf("Key already exists: %+v", key)
		}
	}
	return nil
}

// BatchInsertMap adds all map keys and values to the collection.
// If a key already exists, the value will not be inserted and an error will be returned.
// Use BatchReplaceMap for the alternative functionality.
func (sm *SortedMap) BatchInsertMap(v interface{}) error {
	const unsupportedTypeErr = "Unsupported type."

	switch m := v.(type) {
	case map[interface{}]interface{}:
		return sm.batchInsertMapInterfaceKeys(m)

	case map[string]interface{}:
		return sm.batchInsertMapStringKeys(m)

	default:
		return fmt.Errorf("%s", unsupportedTypeErr)
	}
}
