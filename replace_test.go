package sortedmap

import (
	"testing"

	"github.com/umpc/go-sortedmap/asc"
)

func TestReplace(t *testing.T) {
	records := randRecords(3)
	sm := New(asc.Time)

	for i := 0; i < 5; i++ {
		for ii := range records {
			sm.Replace(records[ii].Key, records[ii].Val)
		}
	}

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestBatchReplace(t *testing.T) {
	if _, _, err := newSortedMapFromRandRecords(1000); err != nil {
		t.Fatal(err)
	}
}