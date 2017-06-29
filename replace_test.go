package sortedmap

import (
	"testing"

	"github.com/umpc/go-sortedmap/asc"
)

func TestReplace(t *testing.T) {
	records := randRecords(3)
	sm := New(asc.Time)

	for i := 0; i < 5; i++ {
		for _, rec := range records {
			sm.Replace(rec.Key, rec.Val)
		}
	}

	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}
}

func TestBatchReplace(t *testing.T) {
	if _, _, err := newSortedMapFromRandRecords(1000); err != nil {
		t.Fatal(err)
	}
}