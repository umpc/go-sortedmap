package sortedmap

func insertString(s []string, i int, str string) []string {
	s = append(s, "")
	copy(s[i + 1:], s[i:])
	s[i] = str

	return s
}

func deleteString(s []string, i int) []string {
	copy(s[i:], s[i + 1:])
	s[len(s) - 1] = ""

	return s[:len(s) - 1]
}