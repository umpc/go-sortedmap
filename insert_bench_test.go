package sortedmap

import (
	"testing"

	"github.com/umpc/go-sortedmap/asc"
)

func BenchmarkBatchInsert1Record(b *testing.B) {
	records := randRecords(1)
	sm := New(asc.Time)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchInsert(records)
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
}