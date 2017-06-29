package sortedmap

import (
	"testing"
	"time"
)

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
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}
}

func TestBatchDelete(t *testing.T) {
	sm, records, err := newSortedMapFromRandRecords(300)
	if err != nil {
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
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteBetween(t *testing.T) {
	const (
		twoNilBoundValsErr = "accepted two nil bound values"
		generalBoundsErr = "only one bound value"
	)
	sm, _, err := newSortedMapFromRandRecords(300)
	if err != nil {
		t.Fatal(err)
	}
	if !sm.DeleteBetween(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()) {
		t.Fatalf("TestDeleteBetween failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}
	if !sm.DeleteBetween(time.Now(), time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)) {
		t.Fatalf("TestDeleteBetween failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}
}