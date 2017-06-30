package sortedmap

import "testing"

func has1ofNRecords(b *testing.B, n int) {
	sm, _, keys, err := newRandSortedMapWithKeys(n)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Has(keys[0])

		b.StopTimer()
		sm, _, keys, err = newRandSortedMapWithKeys(n)
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}

func batchHasNofNRecords(b *testing.B, n int) {
	sm, _, keys, err := newRandSortedMapWithKeys(n)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchHas(keys)

		b.StopTimer()
		sm, _, keys, err = newRandSortedMapWithKeys(n)
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}

func BenchmarkHas1of1CachedRecords(b *testing.B) {
	sm, _, keys, err := newRandSortedMapWithKeys(1)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Has(keys[0])
	}
}

func BenchmarkHas1of1Records(b *testing.B) {
	has1ofNRecords(b, 1)
}

func BenchmarkBatchHas10of10Records(b *testing.B) {
	batchHasNofNRecords(b, 10)
}

func BenchmarkBatchHas100of100Records(b *testing.B) {
	batchHasNofNRecords(b, 100)
}

func BenchmarkBatchHas1000of1000Records(b *testing.B) {
	batchHasNofNRecords(b, 1000)
}

func BenchmarkBatchHas10000of10000Records(b *testing.B) {
	batchHasNofNRecords(b, 10000)
}