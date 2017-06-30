package sortedmap

import (
	"testing"
	"time"
)

const (
	nilBoundValsErr = "accepted nil bound values"
	generalBoundsErr = "between bounds error"
	nilRecErr = "nil record!"
	runawayIterErr = "runaway iterator!"
)

var maxTime = time.Unix(1<<63 - 62135596801, 999999999)

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
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	reversed := false
	earlierDate := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Date(5321, 1, 1, 0, 0, 0, 0, time.UTC)

	_, ok := sm.IterBetweenCh(nil, nil)
	if ok {
		t.Fatalf("TestIterBetweenCh failed: %v", nilBoundValsErr)
	}

	ch, ok := sm.IterBetweenCh(time.Time{}, time.Time{})
	if !ok {
		t.Fatalf("TestIterBetweenCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterBetweenCh(time.Time{}, maxTime)
	if !ok {
		t.Fatalf("TestIterBetweenCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterBetweenCh(maxTime, time.Time{})
	if !ok {
		t.Fatalf("TestIterBetweenCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterBetweenCh(earlierDate, time.Now())
	if !ok {
		t.Fatalf("TestIterBetweenCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterBetweenCh(time.Now(), earlierDate)
	if !ok {
		t.Fatalf("TestIterBetweenCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterBetweenCh(time.Now(), laterDate)
	if !ok {
		t.Fatalf("TestIterBetweenCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterBetweenCh(laterDate, time.Now())
	if !ok {
		t.Fatalf("TestIterBetweenCh failed: %v", generalBoundsErr)
	}

	ch, ok = sm.IterBetweenCh(laterDate, laterDate)
	if !ok {
		t.Fatalf("TestIterBetweenCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}
}

func TestIterBetweenChCustom(t *testing.T) {
	const (
		nilBoundValsErr = "accepted two nil bound values"
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
		t.Fatalf("TestIterBetween failed: %v", nilBoundValsErr)
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
			t.Fatalf("TestIterFunc failed: %v", nilRecErr)
		}
		return true
	})
	sm.IterFunc(true, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestIterFunc failed: %v", nilRecErr)
		}
		return true
	})
	i := 0
	sm.IterFunc(false, func(rec Record) bool {
		if i > 0 {
			t.Fatalf("TestIterFunc failed: %v", runawayIterErr)
		}
		i++
		return false
	})
	i = 0
	sm.IterFunc(true, func(rec Record) bool {
		if i > 0 {
			t.Fatalf("TestIterFunc failed: %v", runawayIterErr)
		}
		i++
		return false
	})
}

func TestIterBetweenFunc(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	earlierDate := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Now()

	sm.IterBetweenFunc(false, nil, nil, func(rec Record) bool {
		t.Fatalf("TestIterBetweenFunc failed: %v", nilBoundValsErr)
		return false
	})

	sm.IterBetweenFunc(false, nil, laterDate, func(rec Record) bool {
		t.Fatalf("TestIterBetweenFunc failed: %v", nilBoundValsErr)
		return false
	})

	sm.IterBetweenFunc(false, laterDate, nil, func(rec Record) bool {
		t.Fatalf("TestIterBetweenFunc failed: %v", nilBoundValsErr)
		return false
	})

	sm.IterBetweenFunc(false, earlierDate, laterDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestIterBetweenFunc failed: %v", nilRecErr)
		}
		return true
	})
	sm.IterBetweenFunc(true, laterDate, earlierDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestIterBetweenFunc failed: %v", nilRecErr)
		}
		return true
	})
	sm.IterBetweenFunc(true, laterDate, earlierDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestIterBetweenFunc failed: %v", nilRecErr)
		}
		return true
	})
	i := 0
	sm.IterBetweenFunc(false, laterDate, earlierDate, func(rec Record) bool {
		if i > 0 {
			t.Fatalf("TestIterBetweenFunc failed: %v", runawayIterErr)
		}
		i++
		return false
	})
	i = 0
	sm.IterBetweenFunc(true, laterDate, earlierDate, func(rec Record) bool {
		if i > 0 {
			t.Fatalf("TestIterBetweenFunc failed: %v", runawayIterErr)
		}
		i++
		return false
	})
}