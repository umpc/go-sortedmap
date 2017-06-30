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

func TestIterChTimeout(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	timeout := 1 * time.Second
	params := &IterChParams{
		SendTimeout: &timeout,
	}
	if _, ok := sm.CustomIterCh(params); ok {
		// Add some form of more advanced timeout test here.
	} else {
		t.Fatalf("TestCustomIterCh failed: %v", generalBoundsErr)
	}
}

func TestBoundedIterCh(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	reversed := false
	earlierDate := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Date(5321, 1, 1, 0, 0, 0, 0, time.UTC)

	_, ok := sm.BoundedIterCh(false, nil, nil)
	if !ok {
		t.Fatalf("TestBoundedIterCh failed: %v", generalBoundsErr)
	}

	ch, ok := sm.BoundedIterCh(false, time.Time{}, time.Time{})
	if !ok {
		t.Fatalf("TestBoundedIterCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.BoundedIterCh(false, time.Time{}, maxTime)
	if !ok {
		t.Fatalf("TestBoundedIterCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.BoundedIterCh(false, maxTime, time.Time{})
	if !ok {
		t.Fatalf("TestBoundedIterCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.BoundedIterCh(false, earlierDate, time.Now())
	if !ok {
		t.Fatalf("TestBoundedIterCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.BoundedIterCh(false, time.Now(), earlierDate)
	if !ok {
		t.Fatalf("TestBoundedIterCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.BoundedIterCh(false, time.Now(), laterDate)
	if !ok {
		t.Fatalf("TestBoundedIterCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}

	ch, ok = sm.BoundedIterCh(false, laterDate, time.Now())
	if !ok {
		t.Fatalf("TestBoundedIterCh failed: %v", generalBoundsErr)
	}

	ch, ok = sm.BoundedIterCh(false, laterDate, laterDate)
	if !ok {
		t.Fatalf("TestBoundedIterCh failed: %v", generalBoundsErr)
	}
	if err := verifyRecords(ch, reversed); err != nil {
		t.Fatal(err)
	}
}

func TestCustomIterCh(t *testing.T) {
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

	params := &IterChParams{
		Reversed: reversed,
		BufSize: buffSize,
	}
	if ch, ok := sm.CustomIterCh(params); ok {
		if err := verifyRecords(ch, params.Reversed); err != nil {
			t.Fatal(err)
		}
	} else {
		t.Fatalf("TestCustomIterCh failed: %v", generalBoundsErr)
	}

	params = &IterChParams{
		Reversed: reversed,
		BufSize: buffSize,
		LowerBound: earlierDate,
		UpperBound: laterDate,
	}
	if ch, ok := sm.CustomIterCh(params); ok {
		if err := verifyRecords(ch, reversed); err != nil {
			t.Fatal(err)
		}
	} else {
		t.Fatalf("TestCustomIterCh failed: %v", generalBoundsErr)
	}

	params = &IterChParams{
		Reversed: reversed,
		BufSize: buffSize,
		LowerBound: laterDate,
		UpperBound: earlierDate,
	}
	if ch, ok := sm.CustomIterCh(params); ok {
		if err := verifyRecords(ch, reversed); err != nil {
			t.Fatal(err)
		}
	} else {
		t.Fatalf("TestCustomIterCh failed: %v", generalBoundsErr)
	}

	reversed = false
	params = &IterChParams{
		Reversed: reversed,
		BufSize: 0,
		LowerBound: laterDate,
		UpperBound: earlierDate,
	}
	if ch, ok := sm.CustomIterCh(params); ok {
		if err := verifyRecords(ch, reversed); err != nil {
			t.Fatal(err)
		}
	} else {
		t.Fatalf("TestCustomIterCh failed: %v", generalBoundsErr)
	}
}

func TestIterFunc(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	sm.IterFunc(false, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestIterFunc failed: %v", nilRecErr)
		}
		return true
	})
	sm.IterFunc(true, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestIterFunc failed: %v", nilRecErr)
		}
		return true
	})
	i := 0
	sm.IterFunc(false, func(rec *Record) bool {
		if i > 0 {
			t.Fatalf("TestIterFunc failed: %v", runawayIterErr)
		}
		i++
		return false
	})
	i = 0
	sm.IterFunc(true, func(rec *Record) bool {
		if i > 0 {
			t.Fatalf("TestIterFunc failed: %v", runawayIterErr)
		}
		i++
		return false
	})
}

func TestBoundedIterFunc(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	earlierDate := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Now()

	sm.BoundedIterFunc(false, nil, nil, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilRecErr)
		}
		return false
	})

	sm.BoundedIterFunc(false, nil, laterDate, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilRecErr)
		}
		return false
	})

	sm.BoundedIterFunc(false, laterDate, nil, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilRecErr)
		}
		return false
	})

	sm.BoundedIterFunc(false, earlierDate, laterDate, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilRecErr)
		}
		return true
	})
	sm.BoundedIterFunc(true, laterDate, earlierDate, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilRecErr)
		}
		return true
	})
	sm.BoundedIterFunc(true, laterDate, earlierDate, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilRecErr)
		}
		return true
	})
	i := 0
	sm.BoundedIterFunc(false, laterDate, earlierDate, func(rec *Record) bool {
		if i > 0 {
			t.Fatalf("TestBoundedIterFunc failed: %v", runawayIterErr)
		}
		i++
		return false
	})
	i = 0
	sm.BoundedIterFunc(true, laterDate, earlierDate, func(rec *Record) bool {
		if i > 0 {
			t.Fatalf("TestBoundedIterFunc failed: %v", runawayIterErr)
		}
		i++
		return false
	})
}