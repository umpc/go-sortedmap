package sortedmap

import (
	"testing"

	"github.com/umpc/go-sortedmap/asc"
)

func TestInsert(t *testing.T) {
	const n = 3
	records := randRecords(n)
	sm := New(n, asc.Time)

	for i := range records {
		if !sm.Insert(records[i].Key, records[i].Val) {
			t.Fatalf("Insert failed: %v", keyExistsErr)
		}
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}

	for i := range records {
		if sm.Insert(records[i].Key, records[i].Val) {
			t.Fatalf("Insert failed: %v", notFoundErr)
		}
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}
}

func TestBatchInsert(t *testing.T) {
	const n = 1000
	records := randRecords(n)
	sm := New(n, asc.Time)

	for _, ok := range sm.BatchInsert(records) {
		if !ok {
			t.Fatalf("BatchInsert failed: %v", keyExistsErr)
		}
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}
}

func TestBatchInsertMapWithInterfaceKeys(t *testing.T) {
	const n = 1000
	records := randRecords(n)
	sm := New(n, asc.Time)

	i := 0
	m := make(map[interface{}]interface{}, n)

	for _, rec := range records {
		m[rec.Key] = rec.Val
		i++
	}
	if i == 0 {
		t.Fatal("Records were not copied to the map.")
	}

	if err := sm.BatchInsertMap(m); err != nil {
		t.Fatal(err)
	}
}

func TestBatchInsertMapWithStringKeys(t *testing.T) {
	const n = 1000
	records := randRecords(n)
	sm := New(n, asc.Time)

	i := 0
	m := make(map[string]interface{}, n)

	for _, rec := range records {
		m[rec.Key.(string)] = rec.Val
		i++
	}
	if i == 0 {
		t.Fatal("Records were not copied to the map.")
	}

	if err := sm.BatchInsertMap(m); err != nil {
		t.Fatal(err)
	}
}