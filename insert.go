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
func (sm *SortedMap) BatchInsert(recs []*Record) []bool {
	results := make([]bool, len(recs))
	for i, rec := range recs {
		results[i] = sm.insert(rec.Key, rec.Val)
	}
	return results
}

func (sm *SortedMap) batchInsertMapWithInterfaceKeys(v interface{}) error {
	m := v.(map[interface{}]interface{})

	for key, val := range m {
		if !sm.insert(key, val) {
			return fmt.Errorf("Key already exists: %+v", key)
		}
	}
	return nil
}

func (sm *SortedMap) batchInsertMapWithStringKeys(v interface{}) error {
	m := v.(map[string]interface{})

	for key, val := range m {
		if !sm.insert(key, val) {
			return fmt.Errorf("Key already exists: %+v", key)
		}
	}
	return nil
}

// BatchInsertMap adds all map keys and values to the collection and returns a slice containing each record's insert status.
// If a key already exists, the value will not be inserted. Use BatchReplaceMap for the alternative functionality.	
func (sm *SortedMap) BatchInsertMap(v interface{}) error {
	const unsupportedTypeErr = "Unsupported type."

	switch v.(type) {
	case map[interface{}]interface{}:
		return sm.batchInsertMapWithInterfaceKeys(v)

	case map[string]interface{}:
		return sm.batchInsertMapWithStringKeys(v)

	default:
		return fmt.Errorf("%s", unsupportedTypeErr)
	}
}