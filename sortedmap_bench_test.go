package sortedmap

import (
	"testing"

	"github.com/umpc/go-sortedmap/asc"
)

func BenchmarkNew(b *testing.B) {
	var sm *SortedMap
	if sm == nil {}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm = New(asc.Time)
	}
}