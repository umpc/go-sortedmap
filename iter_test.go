package sortedmap

import (
	"testing"
	"time"
)

func TestIter(t *testing.T) {
	if _, _, err := newSortedMapFromRandRecords(1000); err != nil {
		t.Fatal(err)
	}
}

func TestIterUntil(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.IterUntil(time.Now()), false); err != nil {
		t.Fatal(err)
	}
}

func TestIterAfter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.IterAfter(time.Now()), false); err != nil {
		t.Fatal(err)
	}
}

func TestReverseIter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.ReverseIter(), true); err != nil {
		t.Fatal(err)
	}
}

func TestReverseIterUntil(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.ReverseIterUntil(time.Now()), true); err != nil {
		t.Fatal(err)
	}
}

func TestReverseIterAfter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.ReverseIterAfter(time.Now()), true); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedIter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedIter(256), false); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedIterUntil(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedIterUntil(256, time.Now()), false); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedIterAfter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedIterAfter(256, time.Now()), false); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedReverseIter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedReverseIter(256), true); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedReverseIterUntil(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedReverseIterUntil(256, time.Now()), true); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedReverseIterAfter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedReverseIterAfter(256, time.Now()), true); err != nil {
		t.Fatal(err)
	}
}