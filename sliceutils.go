package sortedmap

func insertInterface(s []interface{}, v interface{}, i int) []interface{} {
	s = append(s, nil)
	copy(s[i + 1:], s[i:])
	s[i] = v

	return s
}

func deleteInterface(s []interface{}, i int) []interface{} {
	copy(s[i:], s[i + 1:])
	s[len(s) - 1] = nil

	return s[:len(s) - 1]
}