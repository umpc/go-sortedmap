package sortedmap

import (
	"testing"
	"time"
)

const (
	keyExistsErr = "a key already exists in the collection!"
	unsortedErr = "SortedMap is not sorted!"
)

func TestInsert(t *testing.T) {
	records := randRecords(3)
	sm := New(nil)

	for i := range records {
		if !sm.Insert(records[i].Key, records[i].Val) {
			t.Fatalf("Insert failed: %v", keyExistsErr)
		}
	}

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestReplace(t *testing.T) {
	records := randRecords(3)
	sm := New(nil)

	for i := 0; i < 5; i++ {
		for ii := range records {
			sm.Replace(records[ii].Key, records[ii].Val)
		}
	}

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestBatchInsert(t *testing.T) {
	records := randRecords(1000)
	sm := New(nil)

	for _, ok := range sm.BatchInsert(records...) {
		if !ok {
			t.Fatalf("BatchInsert failed: %v", keyExistsErr)
		}
	}

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestBatchReplace(t *testing.T) {
	sm := newSortedMapFromRandRecords(1000)

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestIter(t *testing.T) {
	sm := newSortedMapFromRandRecords(1000)

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestIterUntil(t *testing.T) {
	sm := newSortedMapFromRandRecords(1000)

	if err := verifyRecords(sm.IterUntil(time.Now())); err != nil {
		t.Fatal(err)
	}
}

func TestIterAfter(t *testing.T) {
	sm := newSortedMapFromRandRecords(1000)

	if err := verifyRecords(sm.IterAfter(time.Now())); err != nil {
		t.Fatal(err)
	}
}