package sortedmap

import (
	"testing"
	"time"
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

func TestCancelCustomIterCh(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}

	earlierDate := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	laterDate := time.Now()

	func() {
		params := IterChParams{
			LowerBound: earlierDate,
			UpperBound: laterDate,
		}

		ch, err := sm.CustomIterCh(params)
		if err != nil {
			t.Fatal(err)
		}
		defer ch.Close()

		<-ch.Records()
		ch.Close()

		if err := verifyRecords(ch.Records(), params.Reversed); err != nil {
			if err.Error() != "Channel was nil." {
				t.Fatal(err)
			}
		} else {
			t.Fatal("Channel was not closed.")
		}
	}()
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
