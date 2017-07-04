package sortedmap

// insertInterface inserts the interface{} value v into slice s, at index i.
// and then returns an updated reference.
func insertInterface(s []interface{}, v interface{}, i int) []interface{} {
	s = append(s, nil)
	copy(s[i+1:], s[i:])
	s[i] = v

	return s
}

// deleteInterface deletes an interface{} value from slice s, at index i,
// and then returns an updated reference.
func deleteInterface(s []interface{}, i int) []interface{} {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = nil

	return s[:len(s)-1]
}
