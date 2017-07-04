package sortedmap

import (
	"testing"

	"github.com/umpc/go-sortedmap/asc"
)

func BenchmarkNew(b *testing.B) {
	var sm *SortedMap

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm = New(0, asc.Time)
	}
	b.StopTimer()

	if sm == nil {
	}
}
