package sortedmap

import (
	"testing"

	"github.com/umpc/go-sortedmap/asc"
)

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