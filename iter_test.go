package sortedmap

import (
	"testing"
	"time"
)

func TestIterCh(t *testing.T) {
	if _, _, err := newSortedMapFromRandRecords(1000); err != nil {
		t.Fatal(err)
	}
}

func TestIterChCustom(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	reversed := true
	if err := verifyRecords(sm.IterChCustom(reversed, 256), reversed); err != nil {
		t.Fatal(err)
	}
}

func TestIterBetweenCh(t *testing.T) {
	const (
		twoNilBoundValsErr = "accepted two nil bound values"
		generalBoundsErr = "only one bound value"
	)
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	reversed := false

	_, ok := sm.IterBetweenCh(nil, nil)
	if ok {
		t.Fatalf("TestIterBetween failed: %v", twoNilBoundValsErr)
	}

	ch, ok := sm.IterBetweenCh(time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), time.Now())
	if !ok {
		t.Fatalf("TestIterBetween failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterBetweenCh(time.Now(), time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC))
	if !ok {
		t.Fatalf("TestIterBetween failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}
}

func TestIterBetweenChCustom(t *testing.T) {
	const (
		twoNilBoundValsErr = "accepted two nil bound values"
		generalBoundsErr = "only one bound value"
	)
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	reversed := true
	buffSize := 64

	earlierDate := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Now()

	_, ok := sm.IterBetweenChCustom(reversed, buffSize, nil, nil)
	if ok {
		t.Fatalf("TestIterBetween failed: %v", twoNilBoundValsErr)
	}

	ch, ok := sm.IterBetweenChCustom(reversed, buffSize, earlierDate, laterDate)
	if !ok {
		t.Fatalf("TestIterBetween failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterBetweenChCustom(reversed, buffSize, laterDate, earlierDate)
	if !ok {
		t.Fatalf("TestIterBetween failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	reversed = false
	ch, ok = sm.IterBetweenChCustom(reversed, buffSize, laterDate, earlierDate)
	if !ok {
		t.Fatalf("TestIterBetween failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}
}

func TestIterFunc(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	sm.IterFunc(false, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatal("TestIterFunc failed: nil record!")
		}
		return true
	})
	sm.IterFunc(true, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatal("TestIterFunc failed: nil record!")
		}
		return true
	})
}

func TestIterBetweenFunc(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	earlierDate := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Now()

	sm.IterBetweenFunc(false, earlierDate, laterDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatal("TestIterBetweenFunc failed: nil record!")
		}
		return true
	})
	sm.IterBetweenFunc(true, laterDate, earlierDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatal("TestIterBetweenFunc failed: nil record!")
		}
		return true
	})
}