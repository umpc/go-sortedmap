package sortedmap

import (
	"testing"
	"time"

	"github.com/umpc/go-sortedmap/asc"
)

func BenchmarkNew(b *testing.B) {
	var sm *SortedMap
	if sm == nil {}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm = New(asc.Time)
	}
}

func BenchmarkBatchInsert10Records(b *testing.B) {
	records := randRecords(10)
	sm := New(asc.Time)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchInsert(records)
	}
}

func BenchmarkBatchInsert100Records(b *testing.B) {
	records := randRecords(100)
	sm := New(asc.Time)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchInsert(records)
	}
}

func BenchmarkBatchInsert1000Records(b *testing.B) {
	records := randRecords(1000)
	sm := New(asc.Time)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchInsert(records)
	}
}

func BenchmarkBatchInsert10000Records(b *testing.B) {
	records := randRecords(10000)
	sm := New(asc.Time)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchInsert(records)
	}
	b.StopTimer()

	// Verify
	var previousRec Record
	for rec := range sm.Iter() {
		if previousRec.Key != nil {
			switch previousRec.Val.(type) {
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
	testVal := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)

	records := make([]*Record, 1000)
	for i := range records {
		records[i] = &Record{
			Key: randStr(42),
			Val: testVal,
		}
	}

	sm := New(asc.Time)
	sm.BatchInsert(records)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Has(records[0].Key)
	}
}

func BenchmarkHas10000Records(b *testing.B) {
	testVal := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)

	records := make([]*Record, 10000)
	for i := range records {
		records[i] = &Record{
			Key: randStr(42),
			Val: testVal,
		}
	}

	sm := New(asc.Time)
	sm.BatchInsert(records)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Has(records[0].Key)
	}
}

func BenchmarkDelete1Of1000Records(b *testing.B) {
	testVal := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)

	records := make([]*Record, 1000)
	for i := range records {
		records[i] = &Record{
			Key: randStr(42),
			Val: testVal,
		}
	}

	sm := New(asc.Time)
	sm.BatchInsert(records)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Delete(records[0].Key)
		b.StopTimer()
		sm.BatchInsert([]*Record{records[0]})
		b.StartTimer()
	}
}

func BenchmarkDelete1Of10000Records(b *testing.B) {
	testVal := time.Date(2017, 06, 25, 0, 0, 0, 0, time.UTC)

	records := make([]*Record, 10000)
	for i := range records {
		records[i] = &Record{
			Key: randStr(42),
			Val: testVal,
		}
	}

	sm := New(asc.Time)
	sm.BatchInsert(records)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Delete(records[0].Key)
		b.StopTimer()
		sm.BatchInsert([]*Record{records[0]})
		b.StartTimer()
	}
}