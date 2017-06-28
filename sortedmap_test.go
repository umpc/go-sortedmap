package sortedmap

import "testing"

const (
	notFoundErr   = "key not found!"
	keyExistsErr  = "a key already exists in the collection!"
	unsortedErr   = "SortedMap is not sorted!"
	invalidDelete = "invalid delete status!"
)

func TestNew(t *testing.T) {
	sm := New(nil)

	if sm.idx == nil {
		t.Fatal("TestNew failed: idx was nil!")
	}
	if sm.sorted == nil {
		t.Fatal("TestNew failed: sorted was nil!")
	}
	if sm.lessFn == nil {
		t.Fatal("TestNew failed: lessFn was nil!")
	}
}

func TestFalseLessFunc(t *testing.T) {
	if New(nil).lessFn(nil, nil) {
		t.Fatal("TestFalseLessFunc failed: lessFn returned true!")
	}
}

func TestLen(t *testing.T) {
	const count = 100
	sm, _, err := newSortedMapFromRandRecords(count)
	if err != nil {
		t.Fatal(err)
	}
	if sm.Len() != count {
		t.Fatalf("TestLen failed: invalid SortedMap length. Expected: %v, Had: %v.", count, sm.Len())
	}
}