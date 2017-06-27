package sortedmap

import (
	"testing"
	"time"

	"github.com/umpc/go-sortedmap/asc"
)

const (
	notFoundErr = "key not found!"
	keyExistsErr = "a key already exists in the collection!"
	unsortedErr = "SortedMap is not sorted!"
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

func TestInsert(t *testing.T) {
	records := randRecords(3)
	sm := New(asc.Time)

	for i := range records {
		if !sm.Insert(records[i].Key, records[i].Val) {
			t.Fatalf("Insert failed: %v", keyExistsErr)
		}
	}

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}

	for i := range records {
		if sm.Insert(records[i].Key, records[i].Val) {
			t.Fatalf("Insert failed: %v", notFoundErr)
		}
	}

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

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

func TestHas(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(3)
	if err != nil {
		t.Fatal(err)
	}
	for i := range records {
		if !sm.Has(records[i].Key) {
			t.Fatalf("Has failed: %v", notFoundErr)
		}
	}
	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(3)
	if err != nil {
		t.Fatal(err)
	}
	for i := range records {
		if val, ok := sm.Get(records[i].Key); val == nil || !ok {
			t.Fatalf("Get failed: %v", notFoundErr)
		}
	}
	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(300)
	if err != nil {
		t.Fatal(err)
	}
	if sm.Delete("") {
		t.Fatalf("Delete: %v", invalidDelete)
	}
	for _, rec := range records {
		sm.Delete(rec.Key)
	}
	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestLen(t *testing.T) {
	const count = 100
	sm, _, err := newSortedMapFromRandRecords(count)
	if err != nil {
		t.Fatal(err)
	}
	if sm.Len() != count {
		t.Fatalf("Len: invalid SortedMap length. Expected: %v, Had: %v.", count, sm.Len())
	}
}

func TestBatchInsert(t *testing.T) {
	records := randRecords(1000)
	sm := New(asc.Time)

	for _, ok := range sm.BatchInsert(records) {
		if !ok {
			t.Fatalf("BatchInsert failed: %v", keyExistsErr)
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

func TestBatchHas(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	keys := make([]interface{}, len(records))
	for i := range records {
		keys[i] = records[i].Key
	}

	for _, ok := range sm.BatchHas(keys) {
		if !ok {
			t.Fatalf("BatchHas: %v", notFoundErr)
		}
	}

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestBatchGet(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	keys := make([]interface{}, len(records))
	for i := range records {
		keys[i] = records[i].Key
	}

	values, results := sm.BatchGet(keys)
	for i, ok := range results {
		if values[i] == nil || !ok {
			t.Fatalf("BatchGet: %v", notFoundErr)
		}
	}

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestBatchDelete(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(300)
	if err != nil {
		t.Fatal(err)
	}

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}

	keys := make([]interface{}, 0)
	for i, rec := range records {
		if i == 50 {
			break
		}
		keys = append(keys, rec.Key)
	}

	for _, ok := range sm.BatchDelete(keys) {
		if !ok {
			t.Fatalf("BatchDelete: %v", invalidDelete)
		}
	}

	if err := verifyRecords(sm.Iter()); err != nil {
		t.Fatal(err)
	}
}

func TestIter(t *testing.T) {
	if _, _, err := newSortedMapFromRandRecords(1000); err != nil {
		t.Fatal(err)
	}
}

func TestIterUntil(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.IterUntil(time.Now())); err != nil {
		t.Fatal(err)
	}
}

func TestIterAfter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.IterAfter(time.Now())); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedIter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedIter(256)); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedIterUntil(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedIterUntil(256, time.Now())); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedIterAfter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedIterAfter(256, time.Now())); err != nil {
		t.Fatal(err)
	}
}