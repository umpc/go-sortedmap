package sortedmap

import (
	"fmt"
	"testing"
	"time"
)

func TestBatchInsert(t *testing.T) {
	records := randRecords(1000)
	sm := New(nil)

	results := sm.BatchInsert(records...)
	for i := range results {
		if !results[i] {
			t.Fatal("BatchInsert failed: a key already exists in the collection!")
		}
	}
}

func TestIterAndVerifySortOrder(t *testing.T) {
	records := randRecords(1000)
	sm := New(nil)
	sm.BatchInsert(records...)

	var previousRec Record
	ch := sm.Iter()
	for rec := range ch {
		if previousRec.Key != "" {
			if previousRec.Val.(time.Time).After(rec.Val.(time.Time)) {
				
				t.Fatalf("SortedMap is not sorted! %v\n", fmt.Sprintf("prev: %+v, current: %+v.", previousRec, rec))
			}
		}
		previousRec = rec
	}
}