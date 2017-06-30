package sortedmap

import (
	"testing"
	"time"
)

const (
	nilBoundValsErr = "accepted nil bound values"
	generalBoundsErr = "between bounds error"
	nilValErr = "nil value!"
	nonNilValErr = "non-nil value"
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
	timeout := 1 * time.Microsecond
	sleepDur := 10 * time.Millisecond
	params := &IterChParams{SendTimeout: timeout}

	ch := sm.CustomIterCh(params)
	for i := 0; i < 5; i++ {
		time.Sleep(sleepDur)
		rec := <- ch
		if i > 1 && rec.Key != nil {
			t.Fatalf("TestCustomIterCh failed: %v: %v", nonNilValErr, rec.Key)
		}
	}

	params.LowerBound = time.Time{}
	params.UpperBound = maxTime

	ch = sm.CustomIterCh(params)
	for i := 0; i < 5; i++ {
		time.Sleep(sleepDur)
		rec := <- ch
		if i > 1 && rec.Key != nil {
			t.Fatalf("TestCustomIterCh failed: %v: %v", nonNilValErr, rec.Key)
		}
	}

	params = &IterChParams{
		Reversed: true,
		SendTimeout: timeout,
	}
	ch = sm.CustomIterCh(params)
	for i := 0; i < 5; i++ {
		time.Sleep(sleepDur)
		rec := <- ch
		if i > 1 && rec.Key != nil {
			t.Fatalf("TestCustomIterCh failed: %v: %v", nonNilValErr, rec.Key)
		}
	}

	params.LowerBound = time.Time{}
	params.UpperBound = maxTime

	ch = sm.CustomIterCh(params)
	for i := 0; i < 5; i++ {
		time.Sleep(sleepDur)
		rec := <- ch
		if i > 1 && rec.Key != nil {
			t.Fatalf("TestCustomIterCh failed: %v: %v", nonNilValErr, rec.Key)
		}
	}
}

func TestIterChParamsBounds(t *testing.T) {
	params := new(IterChParams)
	if params.bounds() != nil {
		t.Fatalf("TestIterChParamsBounds failed: %v", nonNilValErr)
	}
	params.LowerBound = 0
	if params.bounds() != nil {
		t.Fatalf("TestIterChParamsBounds failed: %v", nonNilValErr)
	}
	params = new(IterChParams)
	params.UpperBound = 0
	if params.bounds() != nil {
		t.Fatalf("TestIterChParamsBounds failed: %v", nonNilValErr)
	}
	params.LowerBound = 0
	if params.bounds() == nil {
		t.Fatalf("TestIterChParamsBounds failed: %v", nilValErr)
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

	if err := verifyRecords(sm.BoundedIterCh(reversed, nil, nil), reversed); err != nil {
		t.Fatal(err)
	}	

	if err := verifyRecords(sm.BoundedIterCh(reversed, time.Time{}, maxTime), reversed); err != nil {
		t.Fatal(err)
	}

	if err := verifyRecords(sm.BoundedIterCh(reversed, maxTime, time.Time{}), reversed); err != nil {
		t.Fatal(err)
	}

	if err := verifyRecords(sm.BoundedIterCh(reversed, earlierDate, time.Now()), reversed); err != nil {
		t.Fatal(err)
	}

	if err := verifyRecords(sm.BoundedIterCh(reversed, time.Now(), earlierDate), reversed); err != nil {
		t.Fatal(err)
	}

	if err := verifyRecords(sm.BoundedIterCh(reversed, time.Now(), laterDate), reversed); err != nil {
		t.Fatal(err)
	}

	if err := verifyRecords(sm.BoundedIterCh(reversed, laterDate, laterDate), reversed); err != nil {
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
	if err := verifyRecords(sm.CustomIterCh(params), params.Reversed); err != nil {
		t.Fatal(err)
	}

	params = &IterChParams{
		Reversed: reversed,
		BufSize: buffSize,
		LowerBound: earlierDate,
		UpperBound: laterDate,
	}
	if err := verifyRecords(sm.CustomIterCh(params), reversed); err != nil {
		t.Fatal(err)
	}

	params = &IterChParams{
		Reversed: reversed,
		BufSize: buffSize,
		LowerBound: laterDate,
		UpperBound: earlierDate,
	}
	if err := verifyRecords(sm.CustomIterCh(params), reversed); err != nil {
		t.Fatal(err)
	}

	reversed = false
	params = &IterChParams{
		Reversed: reversed,
		BufSize: 0,
		LowerBound: laterDate,
		UpperBound: earlierDate,
	}
	if err := verifyRecords(sm.CustomIterCh(params), reversed); err != nil {
		t.Fatal(err)
	}
}

func TestIterFunc(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	sm.IterFunc(false, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestIterFunc failed: %v", nilValErr)
		}
		return true
	})
	sm.IterFunc(true, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestIterFunc failed: %v", nilValErr)
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
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
		}
		return false
	})

	sm.BoundedIterFunc(false, nil, laterDate, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
		}
		return false
	})

	sm.BoundedIterFunc(false, laterDate, nil, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
		}
		return false
	})

	sm.BoundedIterFunc(false, earlierDate, laterDate, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
		}
		return true
	})
	sm.BoundedIterFunc(true, laterDate, earlierDate, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
		}
		return true
	})
	sm.BoundedIterFunc(true, laterDate, earlierDate, func(rec *Record) bool {
		if rec == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
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