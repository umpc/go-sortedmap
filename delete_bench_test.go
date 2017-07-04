package sortedmap

import "testing"

func delete1ofNRecords(b *testing.B, n int) {
	sm, _, keys, err := newRandSortedMapWithKeys(n)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Delete(keys[0])

		b.StopTimer()
		sm, _, keys, err = newRandSortedMapWithKeys(n)
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}

func batchDeleteNofNRecords(b *testing.B, n int) {
	sm, _, keys, err := newRandSortedMapWithKeys(n)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchDelete(keys)

		b.StopTimer()
		sm, _, keys, err = newRandSortedMapWithKeys(n)
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}

func BenchmarkDelete1of1Records(b *testing.B) {
	delete1ofNRecords(b, 1)
}

func BenchmarkDelete1of10Records(b *testing.B) {
	delete1ofNRecords(b, 10)
}

func BenchmarkDelete1of100Records(b *testing.B) {
	delete1ofNRecords(b, 100)
}

func BenchmarkDelete1of1000Records(b *testing.B) {
	delete1ofNRecords(b, 1000)
}

func BenchmarkDelete1of10000Records(b *testing.B) {
	delete1ofNRecords(b, 10000)
}

func BenchmarkBatchDelete10of10Records(b *testing.B) {
	batchDeleteNofNRecords(b, 10)
}

func BenchmarkBatchDelete100of100Records(b *testing.B) {
	batchDeleteNofNRecords(b, 100)
}

func BenchmarkBatchDelete1000of1000Records(b *testing.B) {
	batchDeleteNofNRecords(b, 1000)
}

func BenchmarkBatchDelete10000of10000Records(b *testing.B) {
	batchDeleteNofNRecords(b, 10000)
}
