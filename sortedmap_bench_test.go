package sortedmap

import (
	"fmt"
	"testing"
	"time"
	mrand "math/rand"
)

func randStr(n int) string {
    mrand.Seed(time.Now().UTC().UnixNano())

	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=~[]{}|:;<>,./?"
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		result[i] = chars[mrand.Intn(len(chars))]
	}
    return string(result)
}

func BenchmarkNew(b *testing.B) {
	var sm *SortedMap
	if sm == nil {}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sm = New(nil)
	}
}

func BenchmarkInsert(b *testing.B) {
	sm := New(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		testJWT := randStr(42)
		testJWTExp := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)
		b.StartTimer()
		sm.Insert(testJWT, testJWTExp)
	}
}

func BenchmarkBatchInsert10Records(b *testing.B) {
	testJWTExp := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)

	records := make([]*Record, 10)
	for i := range records {
		records[i] = &Record{
			Key: randStr(42),
			Val: testJWTExp,
		}
	}

	sm := New(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchInsert(records...)
	}
}

func BenchmarkBatchInsert100Records(b *testing.B) {
	testJWTExp := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)

	records := make([]*Record, 100)
	for i := range records {
		records[i] = &Record{
			Key: randStr(42),
			Val: testJWTExp,
		}
	}

	sm := New(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchInsert(records...)
	}
}

func BenchmarkBatchInsert1000Records(b *testing.B) {
	testJWTExp := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)

	records := make([]*Record, 1000)
	for i := range records {
		records[i] = &Record{
			Key: randStr(42),
			Val: testJWTExp,
		}
	}

	sm := New(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchInsert(records...)
	}
}

func BenchmarkBatchInsert10000Records(b *testing.B) {
	mrand.Seed(time.Now().UTC().UnixNano())

	records := make([]*Record, 10000)
	for i := range records {
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
		t := time.Date(year, mth, day, 0, 0, 0, 0, time.UTC)
		records[i] = &Record{
			Key: t.Format(time.RFC3339),
// 			Val: mrand.Intn(34334534561),
			Val: t,
		}
	}

	sm := New(func(idx map[string]interface{}, sorted []string, i int, val interface{}) bool {
		switch val.(type) {
		case int:
			return val.(int) < idx[sorted[i]].(int)
		case time.Time:
			return val.(time.Time).Before(idx[sorted[i]].(time.Time))
		default:
			return false
		}
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchInsert(records...)
	}
	b.StopTimer()

	// Verify
	var previousRec Record
	for rec := range sm.Iter(-1) {
		if previousRec.Key != "" {
			fmt.Println(previousRec.Val, rec.Val)
			switch previousRec.Val.(type) {
			case int:
				if previousRec.Val.(int) > rec.Val.(int) {
					panic("Sorted map is not sorted!")
				}
			case time.Time:
				if previousRec.Val.(time.Time).After(rec.Val.(time.Time)) {
					panic("Sorted map is not sorted!")
				}
			}
		}
		previousRec = rec
	}
}

func BenchmarkHas1000Records(b *testing.B) {
	testJWTExp := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)

	records := make([]*Record, 1000)
	for i := range records {
		records[i] = &Record{
			Key: randStr(42),
			Val: testJWTExp,
		}
	}

	sm := New(nil)
	sm.BatchInsert(records...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Has(records[0].Key)
	}
}

func BenchmarkHas10000Records(b *testing.B) {
	testJWTExp := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)

	records := make([]*Record, 10000)
	for i := range records {
		records[i] = &Record{
			Key: randStr(42),
			Val: testJWTExp,
		}
	}

	sm := New(nil)
	sm.BatchInsert(records...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Has(records[0].Key)
	}
}

func BenchmarkDelete1Of1000Records(b *testing.B) {
	testJWTExp := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)

	records := make([]*Record, 1000)
	for i := range records {
		records[i] = &Record{
			Key: randStr(42),
			Val: testJWTExp,
		}
	}

	sm := New(nil)
	sm.BatchInsert(records...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Delete(records[0].Key)
		b.StopTimer()
		sm.BatchInsert(records[0])
		b.StartTimer()
	}
}

func BenchmarkDelete1Of10000Records(b *testing.B) {
	testJWTExp := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)

	records := make([]*Record, 10000)
	for i := range records {
		records[i] = &Record{
			Key: randStr(42),
			Val: testJWTExp,
		}
	}

	sm := New(nil)
	sm.BatchInsert(records...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Delete(records[0].Key)
		b.StopTimer()
		sm.BatchInsert(records[0])
		b.StartTimer()
	}
}