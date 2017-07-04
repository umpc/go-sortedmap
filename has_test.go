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

	iterCh := sm.IterCh()
	defer iterCh.Close()

	if err := verifyRecords(iterCh.Records(), false); err != nil {
		t.Fatal(err)
	}
}

func TestBatchHas(t *testing.T) {
	sm, _, keys, err := newRandSortedMapWithKeys(1000)
	if err != nil {
		t.Fatal(err)
	}

	for _, ok := range sm.BatchHas(keys) {
		if !ok {
			t.Fatalf("TestBatchHas failed: %v", notFoundErr)
		}
	}

	iterCh := sm.IterCh()
	defer iterCh.Close()

	if err := verifyRecords(iterCh.Records(), false); err != nil {
		t.Fatal(err)
	}
}
