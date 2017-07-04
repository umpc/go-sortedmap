package sortedmap

import (
	"fmt"
	mrand "math/rand"
	"time"

	"github.com/umpc/go-sortedmap/asc"
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

func randStr(n int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=~[]{}|:;<>,./?"
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		result[i] = chars[mrand.Intn(len(chars))]
	}

	return string(result)
}

func randRecord() Record {
	year := mrand.Intn(2129)
	if year < 1 {
		year++
	}
	mth := time.Month(mrand.Intn(12))
	if mth < 1 {
		mth++
	}
	day := mrand.Intn(28)
	if day < 1 {
		day++
	}
	return Record{
		Key: randStr(42),
		Val: time.Date(year, mth, day, 0, 0, 0, 0, time.UTC),
	}
}

func randRecords(n int) []Record {
	records := make([]Record, n)
	for i := range records {
		records[i] = randRecord()
	}
	return records
}

func verifyRecords(ch <-chan Record, reverse bool) error {
	previousRec := Record{}

	for rec := range ch {
		if previousRec.Key != nil {
			switch reverse {
			case false:
				if previousRec.Val.(time.Time).After(rec.Val.(time.Time)) {
					return fmt.Errorf("%v %v",
						unsortedErr,
						fmt.Sprintf("prev: %+v, current: %+v.", previousRec, rec),
					)
				}
			case true:
				if previousRec.Val.(time.Time).Before(rec.Val.(time.Time)) {
					return fmt.Errorf("%v %v",
						unsortedErr,
						fmt.Sprintf("prev: %+v, current: %+v.", previousRec, rec),
					)
				}
			}
		}
		previousRec = rec
	}

	return nil
}

func newSortedMapFromRandRecords(n int) (*SortedMap, []Record, error) {
	records := randRecords(n)
	sm := New(0, asc.Time)
	sm.BatchReplace(records)

	iterCh := sm.IterCh()
	defer iterCh.Close()

	return sm, records, verifyRecords(iterCh.Records(), false)
}

func newRandSortedMapWithKeys(n int) (*SortedMap, []Record, []interface{}, error) {
	sm, records, err := newSortedMapFromRandRecords(n)
	if err != nil {
		return nil, nil, nil, err
	}
	keys := make([]interface{}, n)
	for n, rec := range records {
		keys[n] = rec.Key
	}
	return sm, records, keys, err
}
