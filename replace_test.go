package sortedmap

import (
	"testing"

	"github.com/umpc/go-sortedmap/asc"
)

func TestReplace(t *testing.T) {
	records := randRecords(3)
	sm := New(0, asc.Time)

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

func TestBatchReplaceMapWithInterfaceKeys(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	m := make(map[interface{}]interface{}, len(records))
	for _, rec := range records {
		m[rec.Key] = rec.Val
		i++
	}
	if i == 0 {
		t.Fatal("Records were not copied to the map.")
	}
	if err := sm.BatchReplaceMap(m); err != nil {
		t.Fatal(err)
	}
}

func TestBatchReplaceMapWithStringKeys(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	m := make(map[string]interface{}, len(records))
	for _, rec := range records {
		m[rec.Key.(string)] = rec.Val
		i++
	}
	if i == 0 {
		t.Fatal("Records were not copied to the map.")
	}
	if err := sm.BatchReplaceMap(m); err != nil {
		t.Fatal(err)
	}
}

func TestBatchReplaceMapWithNilType(t *testing.T) {
	if err := New(0, asc.Time).BatchReplaceMap(nil); err == nil {
		t.Fatal("a nil type was allowed where a supported map type is required.")
	}
}