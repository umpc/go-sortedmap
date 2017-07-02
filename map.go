package sortedmap

// Map returns a map containing keys mapped to values.
// The returned map is valid until the next modification to the SortedMap structure.
// The map can be used with ether the Keys or BoundedKeys methods to select a range of items
// and iterate over them using a slice for-range loop, rather than a channel for-range loop.
func (sm *SortedMap) Map() map[interface{}]interface{} {
	return sm.idx
}