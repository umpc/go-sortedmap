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

func TestBoundedDelete(t *testing.T) {
	const (
		nilBoundValsErr = "accepted nil bound value"
		generalBoundsErr = "general bounds error"
	)

	sm, _, err := newSortedMapFromRandRecords(300)
	if err != nil {
		t.Fatal(err)
	}

	earlierDate := time.Date(200, 1, 1, 0, 0, 0, 0, time.UTC)

	if !sm.BoundedDelete(nil, nil) {
		t.Fatalf("TestBoundedDelete failed: %v", generalBoundsErr)
	}

	if !sm.BoundedDelete(nil, time.Now()) {
		t.Fatalf("TestBoundedDelete failed: %v", generalBoundsErr)
	}

	if !sm.BoundedDelete(time.Now(), nil) {
		t.Fatalf("TestBoundedDelete failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}

	if !sm.BoundedDelete(earlierDate, time.Now()) {
		t.Fatalf("TestBoundedDelete failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}

	if !sm.BoundedDelete(time.Now(), earlierDate) {
		t.Fatalf("TestBoundedDelete failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}

	if sm.BoundedDelete(earlierDate, earlierDate) {
		t.Fatalf("TestBoundedDelete failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}
}