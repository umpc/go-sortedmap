package sortedmap

import "testing"

func get1ofNRecords(b *testing.B, n int) {
	sm, _, keys, err := newRandSortedMapWithKeys(n)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Get(keys[0])

		b.StopTimer()
		sm, _, keys, err = newRandSortedMapWithKeys(n)
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}

func batchGetNofNRecords(b *testing.B, n int) {
	sm, _, keys, err := newRandSortedMapWithKeys(n)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchGet(keys)

		b.StopTimer()
		sm, _, keys, err = newRandSortedMapWithKeys(n)
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}

func BenchmarkGet1of1CachedRecords(b *testing.B) {
	sm, _, keys, err := newRandSortedMapWithKeys(1)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Get(keys[0])
	}
}

func BenchmarkGet1of1Records(b *testing.B) {
	get1ofNRecords(b, 1)
}

func BenchmarkBatchGet10of10Records(b *testing.B) {
	batchGetNofNRecords(b, 10)
}

func BenchmarkBatchGet100of100Records(b *testing.B) {
	batchGetNofNRecords(b, 100)
}

func BenchmarkBatchGet1000of1000Records(b *testing.B) {
	batchGetNofNRecords(b, 1000)
}

func BenchmarkBatchGet10000of10000Records(b *testing.B) {
	batchGetNofNRecords(b, 10000)
}