package sortedmap

import "testing"

func BenchmarkBatchReplace1of1Record(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(1)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchReplace(records)
	}
}

func BenchmarkBatchReplace10of10Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(10)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchReplace(records)
	}
}

func BenchmarkBatchReplace100of100Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(100)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchReplace(records)
	}
}

func BenchmarkBatchReplace1000of1000Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchReplace(records)
	}
}

func BenchmarkBatchReplace10000of10000Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(10000)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchReplace(records)
	}
}