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

func TestIterRangeCh(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	reversed := false
	earlierDate := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Date(5321, 1, 1, 0, 0, 0, 0, time.UTC)

	_, ok := sm.IterRangeCh(nil, nil)
	if ok {
		t.Fatalf("TestIterRangeCh failed: %v", nilBoundValsErr)
	}

	ch, ok := sm.IterRangeCh(time.Time{}, time.Time{})
	if !ok {
		t.Fatalf("TestIterRangeCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterRangeCh(time.Time{}, maxTime)
	if !ok {
		t.Fatalf("TestIterRangeCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterRangeCh(maxTime, time.Time{})
	if !ok {
		t.Fatalf("TestIterRangeCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterRangeCh(earlierDate, time.Now())
	if !ok {
		t.Fatalf("TestIterRangeCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterRangeCh(time.Now(), earlierDate)
	if !ok {
		t.Fatalf("TestIterRangeCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterRangeCh(time.Now(), laterDate)
	if !ok {
		t.Fatalf("TestIterRangeCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterRangeCh(laterDate, time.Now())
	if !ok {
		t.Fatalf("TestIterRangeCh failed: %v", generalBoundsErr)
	}

	ch, ok = sm.IterRangeCh(laterDate, laterDate)
	if !ok {
		t.Fatalf("TestIterRangeCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}
}

func TestIterRangeChCustom(t *testing.T) {
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

	_, ok := sm.IterRangeChCustom(reversed, buffSize, nil, nil)
	if ok {
		t.Fatalf("TestIterRange failed: %v", nilBoundValsErr)
	}

	ch, ok := sm.IterRangeChCustom(reversed, buffSize, earlierDate, laterDate)
	if !ok {
		t.Fatalf("TestIterRange failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.IterRangeChCustom(reversed, buffSize, laterDate, earlierDate)
	if !ok {
		t.Fatalf("TestIterRange failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	reversed = false
	ch, ok = sm.IterRangeChCustom(reversed, buffSize, laterDate, earlierDate)
	if !ok {
		t.Fatalf("TestIterRange failed: %v", generalBoundsErr)
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

func TestIterRangeFunc(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	earlierDate := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Now()

	sm.IterRangeFunc(false, nil, nil, func(rec Record) bool {
		t.Fatalf("TestIterRangeFunc failed: %v", nilBoundValsErr)
		return false
	})

	sm.IterRangeFunc(false, nil, laterDate, func(rec Record) bool {
		t.Fatalf("TestIterRangeFunc failed: %v", nilBoundValsErr)
		return false
	})

	sm.IterRangeFunc(false, laterDate, nil, func(rec Record) bool {
		t.Fatalf("TestIterRangeFunc failed: %v", nilBoundValsErr)
		return false
	})

	sm.IterRangeFunc(false, earlierDate, laterDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestIterRangeFunc failed: %v", nilRecErr)
		}
		return true
	})
	sm.IterRangeFunc(true, laterDate, earlierDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestIterRangeFunc failed: %v", nilRecErr)
		}
		return true
	})
	sm.IterRangeFunc(true, laterDate, earlierDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestIterRangeFunc failed: %v", nilRecErr)
		}
		return true
	})
	i := 0
	sm.IterRangeFunc(false, laterDate, earlierDate, func(rec Record) bool {
		if i > 0 {
			t.Fatalf("TestIterRangeFunc failed: %v", runawayIterErr)
		}
		i++
		return false
	})
	i = 0
	sm.IterRangeFunc(true, laterDate, earlierDate, func(rec Record) bool {
		if i > 0 {
			t.Fatalf("TestIterRangeFunc failed: %v", runawayIterErr)
		}
		i++
		return false
	})
}