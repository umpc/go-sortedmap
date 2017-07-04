package sortedmap

import "errors"

func (sm *SortedMap) keys(lowerBound, upperBound interface{}) ([]interface{}, error) {
	idxBounds := sm.boundsIdxSearch(lowerBound, upperBound)
	if idxBounds == nil {
		return nil, errors.New(noValuesErr)
	}
	return sm.sorted[idxBounds[0] : idxBounds[1]+1], nil
}

// Keys returns a slice containing sorted keys.
// The returned slice is valid until the next modification to the SortedMap structure.
func (sm *SortedMap) Keys() []interface{} {
	keys, _ := sm.keys(nil, nil)
	return keys
}

// BoundedKeys returns a slice containing sorted keys equal to or between the given bounds.
// The returned slice is valid until the next modification to the SortedMap structure.
func (sm *SortedMap) BoundedKeys(lowerBound, upperBound interface{}) ([]interface{}, error) {
	return sm.keys(lowerBound, upperBound)
}
