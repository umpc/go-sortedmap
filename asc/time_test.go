package asc

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	earlerDate := time.Date(2017, 06, 05, 18, 0, 0, 0, time.UTC)
	laterDate := time.Date(2018, 07, 06, 21, 0, 0, 0, time.UTC)

	if Time(laterDate, earlerDate) {
		t.Fatal("asc.TestTime failed: laterDate was before earlierDate")
	}
}