package sortedmap

import (
	"testing"
	"time"
)

func TestKeys(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(300)
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	keys := sm.Keys()
	for _, key := range keys {
		if key == nil {
			t.Fatal("Key's value is nil.")
		}
		i++
	}
	if i == 0 {
		t.Fatal("The returned slice was empty.")
	}
}

func TestBoundedKeys(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(300)
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	keys, err := sm.BoundedKeys(time.Time{}, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range keys {
		if key == nil {
			t.Fatal("Key's value is nil.")
		}
		i++
	}
	if i == 0 {
		t.Fatal("The returned slice was empty.")
	}
}

func TestBoundedKeysWithNoBoundsReturned(t *testing.T) {
	sm, _, err := newSortedMapFromRandRecords(300)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := sm.BoundedKeys(time.Now().Add(-1*time.Microsecond), time.Now()); err != nil {
		t.Fatal(err)
	}
}
