package sortedmap

import "testing"

func BenchmarkDelete1of1Record(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(1)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Delete(records[0].Key)

		b.StopTimer()
		sm.Insert(records[0].Key, records[0].Val)
		b.StartTimer()
	}
}

func BenchmarkDelete1of10Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(10)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Delete(records[0].Key)

		b.StopTimer()
		sm.Insert(records[0].Key, records[0].Val)
		b.StartTimer()
	}
}

func BenchmarkDelete1of100Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(100)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Delete(records[0].Key)

		b.StopTimer()
		sm.Insert(records[0].Key, records[0].Val)
		b.StartTimer()
	}
}

func BenchmarkDelete1of1000Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(1000)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Delete(records[0].Key)

		b.StopTimer()
		sm.Insert(records[0].Key, records[0].Val)
		b.StartTimer()
	}
}

func BenchmarkDelete1of10000Records(b *testing.B) {
	sm, records, err := newSortedMapFromRandRecords(10000)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Delete(records[0].Key)

		b.StopTimer()
		sm.Insert(records[0].Key, records[0].Val)
		b.StartTimer()
	}
}