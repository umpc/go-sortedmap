package sortedmap

import "testing"

func TestHas(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(3)
	if err != nil {
		t.Fatal(err)
	}
	for i := range records {
		if !sm.Has(records[i].Key) {
			t.Fatalf("TestHas failed: %v", notFoundErr)
		}
	}
	if err := verifyRecords(sm.Iter()); err != nil {
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