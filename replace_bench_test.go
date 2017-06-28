package sortedmap

import "testing"

func replace1ofNRecords(b *testing.B, n int) {
	sm, records, _, err := newRandSortedMapWithKeys(n)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Replace(records[0].Key, records[0].Val)

		b.StopTimer()
		sm, records, _, err = newRandSortedMapWithKeys(n)
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}

func batchReplaceNofNRecords(b *testing.B, n int) {
	sm, records, _, err := newRandSortedMapWithKeys(n)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.BatchReplace(records)

		b.StopTimer()
		sm, records, _, err = newRandSortedMapWithKeys(n)
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}

func BenchmarkReplace1of1Records(b *testing.B) {
	replace1ofNRecords(b, 1)
}

func BenchmarkReplace1of10Records(b *testing.B) {
	replace1ofNRecords(b, 10)
}

func BenchmarkReplace1of100Records(b *testing.B) {
	replace1ofNRecords(b, 100)
}

func BenchmarkReplace1of1000Records(b *testing.B) {
	replace1ofNRecords(b, 1000)
}

func BenchmarkReplace1of10000Records(b *testing.B) {
	replace1ofNRecords(b, 10000)
}

func BenchmarkBatchReplace10of10Records(b *testing.B) {
	batchDeleteNofNRecords(b, 10)
}

func BenchmarkBatchReplace100of100Records(b *testing.B) {
	batchDeleteNofNRecords(b, 100)
}

func BenchmarkBatchReplace1000of1000Records(b *testing.B) {
	batchDeleteNofNRecords(b, 1000)
}

func BenchmarkBatchReplace10000of10000Records(b *testing.B) {
	batchDeleteNofNRecords(b, 10000)
}