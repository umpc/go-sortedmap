package sortedmap

import (
	"testing"

	"github.com/umpc/go-sortedmap/asc"
)

func insertRecord(b *testing.B) {
	records := randRecords(1)
	sm := New(0, asc.Time)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Insert(records[0].Key, records[0].Val)

		b.StopTimer()
		records = randRecords(1)
		sm = New(0, asc.Time)
		b.StartTimer()
	}
}

func batchInsertRecords(b *testing.B, n int) {
	records := randRecords(n)
	sm := New(0, asc.Time)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchInsert(records)

		b.StopTimer()
		records = randRecords(n)
		sm = New(0, asc.Time)
		b.StartTimer()
	}
}

func BenchmarkInsert1Record(b *testing.B) {
	insertRecord(b)
}

func BenchmarkBatchInsert10Records(b *testing.B) {
	batchInsertRecords(b, 10)
}

func BenchmarkBatchInsert100Records(b *testing.B) {
	batchInsertRecords(b, 100)
}

func BenchmarkBatchInsert1000Records(b *testing.B) {
	batchInsertRecords(b, 1000)
}

func BenchmarkBatchInsert10000Records(b *testing.B) {
	batchInsertRecords(b, 10000)
}
