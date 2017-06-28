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
	if err := verifyRecords(sm.IterUntil(time.Now())); err != nil {
		t.Fatal(err)
	}
}

func TestIterAfter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.IterAfter(time.Now())); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedIter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedIter(256)); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedIterUntil(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedIterUntil(256, time.Now())); err != nil {
		t.Fatal(err)
	}
}

func TestBufferedIterAfter(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyRecords(sm.BufferedIterAfter(256, time.Now())); err != nil {
		t.Fatal(err)
	}
}