package sortedmap

import "testing"

func BenchmarkGet1Record(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(1)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Get(records[0].Key)
	}
}

func BenchmarkGet10Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(10)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Get(records[0].Key)
	}
}

func BenchmarkGet100Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(100)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Get(records[0].Key)
	}
}

func BenchmarkGet1000Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Get(records[0].Key)
	}
}

func BenchmarkGet10000Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(10000)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Has(records[0].Key)
	}
}
