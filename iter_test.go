package sortedmap

import (
	"testing"
	"time"

	"github.com/umpc/go-sortedmap/asc"
)

const (
	nilBoundValsErr  = "accepted nil bound values"
	generalBoundsErr = "between bounds error"
	nilValErr        = "nil value!"
	nonNilValErr     = "non-nil value"
	runawayIterErr   = "runaway iterator!"
)

var maxTime = time.Unix(1<<63-62135596801, 999999999)

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

	params := IterChParams{
		SendTimeout: timeout,
	}

	ch, err := sm.CustomIterCh(params)
	if err != nil {
		t.Fatal(err)
	} else {
		func() {
			defer ch.Close()

			for i := 0; i < 5; i++ {
				time.Sleep(sleepDur)
				rec := <-ch.Records()
				if i > 1 && rec.Key != nil {
					t.Fatalf("TestIterChTimeout failed: %v: %v", nonNilValErr, rec.Key)
				}
			}
		}()
	}

	params.LowerBound = time.Time{}
	params.UpperBound = maxTime

	ch, err = sm.CustomIterCh(params)
	if err != nil {
		t.Fatal(err)
	} else {
		func() {
			defer ch.Close()

			for i := 0; i < 5; i++ {
				time.Sleep(sleepDur)
				rec := <-ch.Records()
				if i > 1 && rec.Key != nil {
					t.Fatalf("TestIterChTimeout failed: %v: %v", nonNilValErr, rec.Key)
				}
			}
		}()
	}
}

func TestReversedIterChTimeout(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	timeout := 1 * time.Microsecond
	sleepDur := 10 * time.Millisecond

	params := IterChParams{
		Reversed:    true,
		SendTimeout: timeout,
	}

	ch, err := sm.CustomIterCh(params)
	if err != nil {
		t.Fatal(err)
	} else {
		func() {
			defer ch.Close()

			for i := 0; i < 5; i++ {
				time.Sleep(sleepDur)
				rec := <-ch.Records()
				if i > 1 && rec.Key != nil {
					t.Fatalf("TestReversedIterChTimeout failed: %v: %v", nonNilValErr, rec.Key)
				}
			}
		}()
	}

	params.LowerBound = time.Time{}
	params.UpperBound = maxTime

	ch, err = sm.CustomIterCh(params)
	if err != nil {
		t.Fatal(err)
	} else {
		func() {
			defer ch.Close()

			for i := 0; i < 5; i++ {
				time.Sleep(sleepDur)
				rec := <-ch.Records()
				if i > 1 && rec.Key != nil {
					t.Fatalf("TestReversedIterChTimeout failed: %v: %v", nonNilValErr, rec.Key)
				}
			}
		}()
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

	ch, err := sm.BoundedIterCh(reversed, nil, nil)
	if err != nil {
		t.Fatal(err)
	} else {
		func() {
			defer ch.Close()

			if err := verifyRecords(ch.Records(), reversed); err != nil {
				t.Fatal(err)
			}
		}()
	}

	ch, err = sm.BoundedIterCh(reversed, time.Time{}, maxTime)
	if err != nil {
		t.Fatal(err)
	} else {
		func() {
			defer ch.Close()

			if err := verifyRecords(ch.Records(), reversed); err != nil {
				t.Fatal(err)
			}
		}()
	}

	ch, err = sm.BoundedIterCh(reversed, earlierDate, time.Now())
	if err != nil {
		t.Fatal(err)
	} else {
		func() {
			defer ch.Close()

			if err := verifyRecords(ch.Records(), reversed); err != nil {
				t.Fatal(err)
			}
		}()
	}

	ch, err = sm.BoundedIterCh(reversed, time.Now(), laterDate)
	if err != nil {
		t.Fatal(err)
	} else {
		func() {
			defer ch.Close()

			if err := verifyRecords(ch.Records(), reversed); err != nil {
				t.Fatal(err)
			}
		}()
	}

	if _, err := sm.BoundedIterCh(reversed, laterDate, laterDate); err == nil {
		t.Fatalf("TestBoundedIterCh failed: %v", "equal bounds values were accepted error")
	}
}

func TestBounds(t *testing.T) {
	sm := New(4, asc.Time)

	obsd := time.Date(1995, 10, 18, 8, 37, 1, 0, time.UTC)
	unixtime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	linux := time.Date(1991, 8, 25, 20, 57, 8, 0, time.UTC)
	github := time.Date(2008, 4, 10, 0, 0, 0, 0, time.UTC)

	sm.Insert("OpenBSD", obsd)
	reversed := false

	// Test bounds combinations which should not match against the value currently in the map
	for _, bounds := range [][]interface{}{
		{unixtime, linux}, {nil, unixtime}, {github, nil},
	} {
		_, err := sm.BoundedIterCh(reversed, bounds[0], bounds[1])
		if err == nil || err.Error() != noValuesErr {
			t.Fatalf("expected no values match error using bounds (lower: %v, upper: %v)", bounds[0], bounds[1])
		}
	}

	// Test bounds combinations which should match against the value currently in the map
	for _, bounds := range [][]interface{}{
		{unixtime, github}, {unixtime, nil}, {nil, github},
	} {
		iterCh, err := sm.BoundedIterCh(reversed, bounds[0], bounds[1])
		if err != nil {
			t.Fatal(err)
		}
		for rec := range iterCh.Records() {
			if rec.Val.(time.Time) != obsd {
				t.Fatal("unexpected value returned by bounded iterator")
			}
		}
		iterCh.Close()
	}

	sm.Insert("UnixTime", unixtime)
	sm.Insert("Linux", linux)
	sm.Insert("GitHub", github)

	func() {
		iterCh, err := sm.BoundedIterCh(reversed, time.Time{}, unixtime)
		if err != nil {
			t.Fatal(err)
		} else {
			defer iterCh.Close()

			if err := verifyRecords(iterCh.Records(), reversed); err != nil {
				t.Fatal(err)
			}
		}
	}()

	func() {
		iterCh, err := sm.BoundedIterCh(reversed, obsd, github)
		if err != nil {
			t.Fatal(err)
		} else {
			defer iterCh.Close()

			if err := verifyRecords(iterCh.Records(), reversed); err != nil {
				t.Fatal(err)
			}
		}
	}()

	_, err := sm.BoundedIterCh(reversed, obsd, obsd)
	if err == nil {
		t.Fatal("equal bounds values were accepted error")
	}
}

func TestCustomIterCh(t *testing.T) {
	const (
		nilBoundValsErr  = "accepted two nil bound values"
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

	params := IterChParams{
		Reversed: reversed,
		BufSize:  buffSize,
	}

	func() {
		ch, err := sm.CustomIterCh(params)
		if err != nil {
			t.Fatal(err)
		}
		defer ch.Close()

		if err := verifyRecords(ch.Records(), params.Reversed); err != nil {
			t.Fatal(err)
		}
	}()

	params = IterChParams{
		Reversed:   reversed,
		BufSize:    buffSize,
		LowerBound: earlierDate,
		UpperBound: laterDate,
	}

	func() {
		ch, err := sm.CustomIterCh(params)
		if err != nil {
			t.Fatal(err)
		}
		defer ch.Close()

		if err := verifyRecords(ch.Records(), params.Reversed); err != nil {
			t.Fatal(err)
		}
	}()

	params = IterChParams{
		Reversed:   reversed,
		BufSize:    buffSize,
		LowerBound: laterDate,
		UpperBound: earlierDate,
	}

	func() {
		_, err := sm.CustomIterCh(params)
		if err == nil {
			t.Fatal(err)
		}
	}()

	reversed = false
	params = IterChParams{
		Reversed:   reversed,
		BufSize:    0,
		LowerBound: laterDate,
		UpperBound: earlierDate,
	}

	func() {
		_, err := sm.CustomIterCh(params)
		if err == nil {
			t.Fatal(err)
		}
	}()
}

func TestCloseCustomIterCh(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	earlierDate := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Now()

	ch, err := sm.IterCh()
	if err != nil {
		t.Fatal(err)
	}

	ch.Close()

	params := IterChParams{
		SendTimeout: 5 * time.Minute,
		LowerBound:  earlierDate,
		UpperBound:  laterDate,
	}

	ch, err = sm.CustomIterCh(params)
	if err != nil {
		t.Fatal(err)
	}

	ch.Close()
}

func TestIterFunc(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	sm.IterFunc(false, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestIterFunc failed: %v", nilValErr)
		}
		return true
	})
	sm.IterFunc(true, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestIterFunc failed: %v", nilValErr)
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

func TestBoundedIterFunc(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	earlierDate := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Now()

	if err := sm.BoundedIterFunc(false, nil, nil, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
		}
		return false
	}); err != nil {
		t.Fatalf("TestBoundedIterFunc failed: %v", err)
	}

	if err := sm.BoundedIterFunc(false, nil, laterDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
		}
		return false
	}); err != nil {
		t.Fatalf("TestBoundedIterFunc failed: %v", err)
	}

	if err := sm.BoundedIterFunc(false, laterDate, nil, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
		}
		return false
	}); err != nil {
		t.Fatalf("TestBoundedIterFunc failed: %v", err)
	}

	if err := sm.BoundedIterFunc(false, earlierDate, laterDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
		}
		return false
	}); err != nil {
		t.Fatalf("TestBoundedIterFunc failed: %v", err)
	}
}

func TestBoundedIterFuncWithNoBoundsReturned(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := sm.BoundedIterFunc(false, time.Date(5783, 1, 1, 0, 0, 0, 0, time.UTC), time.Now(), func(rec Record) bool {
		if rec.Key == nil {
			t.Fatal(nilValErr)
		}
		return false
	}); err == nil {
		t.Fatal("Values fall between or are equal to the given bounds when it should not have returned bounds.")
	}
}

func TestReversedBoundedIterFunc(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	earlierDate := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Now()

	if err := sm.BoundedIterFunc(true, nil, nil, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestReversedBoundedIterFunc failed: %v", nilValErr)
		}
		return false
	}); err != nil {
		t.Fatalf("TestReversedBoundedIterFunc failed: %v", err)
	}

	if err := sm.BoundedIterFunc(true, nil, laterDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestReversedBoundedIterFunc failed: %v", nilValErr)
		}
		return false
	}); err != nil {
		t.Fatalf("TestReversedBoundedIterFunc failed: %v", err)
	}

	if err := sm.BoundedIterFunc(true, laterDate, nil, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
		}
		return false
	}); err != nil {
		t.Fatalf("TestReversedBoundedIterFunc failed: %v", err)
	}

	if err := sm.BoundedIterFunc(true, earlierDate, laterDate, func(rec Record) bool {
		if rec.Key == nil {
			t.Fatalf("TestBoundedIterFunc failed: %v", nilValErr)
		}
		return false
	}); err != nil {
		t.Fatalf("TestReversedBoundedIterFunc failed: %v", err)
	}
}
