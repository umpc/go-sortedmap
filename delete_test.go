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

	iterCh := sm.IterCh()
	if err != nil {
		t.Fatal(err)
	}
	defer iterCh.Close()

	if err := verifyRecords(iterCh.Records(), false); err != nil {
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

	iterCh := sm.IterCh()
	defer iterCh.Close()

	if err := verifyRecords(iterCh.Records(), false); err != nil {
		t.Fatal(err)
	}
}

func TestBoundedDelete(t *testing.T) {
	const (
		nilBoundValsErr  = "accepted nil bound value"
		generalBoundsErr = "general bounds error"
	)

	sm, _, err := newSortedMapFromRandRecords(300)
	if err != nil {
		t.Fatal(err)
	}

	earlierDate := time.Date(200, 1, 1, 0, 0, 0, 0, time.UTC)

	if err := sm.BoundedDelete(nil, nil); err != nil {
		t.Fatal(err)
	}

	if err := sm.BoundedDelete(nil, time.Now()); err != nil {
		t.Fatal(err)
	}

	if err := sm.BoundedDelete(time.Now(), nil); err != nil {
		t.Fatal(err)
	}

	func() {
		iterCh := sm.IterCh()
		defer iterCh.Close()

		if err := verifyRecords(iterCh.Records(), false); err != nil {
			t.Fatal(err)
		}
	}()

	if err := sm.BoundedDelete(earlierDate, time.Now()); err != nil {
		t.Fatal(err)
	}

	func() {
		iterCh := sm.IterCh()
		defer iterCh.Close()

		if err := verifyRecords(iterCh.Records(), false); err != nil {
			t.Fatal(err)
		}
	}()

	if err := sm.BoundedDelete(time.Now(), earlierDate); err == nil {
		t.Fatal(err)
	}

	func() {
		iterCh := sm.IterCh()
		defer iterCh.Close()

		if err := verifyRecords(iterCh.Records(), false); err != nil {
			t.Fatal(err)
		}
	}()

	if err := sm.BoundedDelete(earlierDate, earlierDate); err != nil {
		t.Fatal(err)
	}

	func() {
		iterCh := sm.IterCh()
		defer iterCh.Close()

		if err := verifyRecords(iterCh.Records(), false); err != nil {
			t.Fatal(err)
		}
	}()
}
