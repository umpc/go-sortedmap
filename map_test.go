package sortedmap

import "testing"

func TestMap(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(300)
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	m := sm.Map()
	for _, val := range m {
		if val == nil {
			t.Fatal("Map key's value is nil.")
		}
		i++
	}
	if i == 0 {
		t.Fatal("The returned map was empty.")
	}
}
