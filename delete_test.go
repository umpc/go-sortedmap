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

func TestDeleteRange(t *testing.T) {
	const (
		nilBoundValsErr = "accepted nil bound value"
		generalBoundsErr = "general bounds error"
	)

	sm, _, err := newSortedMapFromRandRecords(300)
	if err != nil {
		t.Fatal(err)
	}

	earlierDate := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

	if sm.DeleteRange(nil, nil) {
		t.Fatalf("TestDeleteRange failed: %v", generalBoundsErr)
	}

	if sm.DeleteRange(nil, time.Now()) {
		t.Fatalf("TestDeleteRange failed: %v", generalBoundsErr)
	}

	if sm.DeleteRange(time.Now(), nil) {
		t.Fatalf("TestDeleteRange failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}

	if !sm.DeleteRange(earlierDate, time.Now()) {
		t.Fatalf("TestDeleteRange failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}

	if !sm.DeleteRange(time.Now(), earlierDate) {
		t.Fatalf("TestDeleteRange failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}

	if !sm.DeleteRange(earlierDate, earlierDate) {
		t.Fatalf("TestDeleteRange failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(sm.IterCh(), false); err != nil {
		t.Fatal(err)
	}
}