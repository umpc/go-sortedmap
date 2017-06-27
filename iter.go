package sortedmap

import "sort"

func (sm *SortedMap) iter(bufSize int) <-chan Record {
	ch := make(chan Record, bufSize)

	go func(ch chan Record) {
		var (
			key string
			idxSnapshot = sm.idx
			sortedSnapshot = sm.sorted
			smLen = len(sortedSnapshot)
		)
		for i := 0; i < smLen; i++ {
			key = sortedSnapshot[i]
			ch <- Record{
				Key: key,
				Val: idxSnapshot[key],
			}
		}
		close(ch)
	}(ch)

	return ch
}

func (sm *SortedMap) iterUntil(bufSize int, val interface{}) <-chan Record {
	ch := make(chan Record, bufSize)

	go func(ch chan Record) {
		var (
			key string
			idxSnapshot = sm.idx
			sortedSnapshot = sm.sorted
			smLen = len(sortedSnapshot)
		)
		for i := 0; i < smLen; i++ {
			if sm.lessFn(idxSnapshot, sortedSnapshot, i, val) {
				break
			}
			key = sortedSnapshot[i]
			ch <- Record{
				Key: key,
				Val: idxSnapshot[key],
			}
		}
		close(ch)
	}(ch)

	return ch
}

func (sm *SortedMap) iterAfter(bufSize int, val interface{}) <-chan Record {
	ch := make(chan Record, bufSize)

	go func(ch chan Record) {
		var (
			key string
			idxSnapshot = sm.idx
			sortedSnapshot = sm.sorted
			smLen = len(sortedSnapshot)
		)
		i := sort.Search(smLen, func(i int) bool {
			return sm.lessFn(idxSnapshot, sortedSnapshot, i, val)
		})
		for; i < smLen; i++ {
			key = sortedSnapshot[i]
			ch <- Record{
				Key: key,
				Val: idxSnapshot[key],
			}
		}
		close(ch)
	}(ch)

	return ch
}

// Iter returns an unbuffered channel that sorted records can be read from and processed.
func (sm *SortedMap) Iter() <-chan Record {
	return sm.iter(0)
}

// IterUntil returns an unbuffered channel that sorted records can be read from and processed.
// IterUntil starts at the lowest value in the collection and sends all values until reaching the given value.
func (sm *SortedMap) IterUntil(val interface{}) <-chan Record {
	return sm.iterUntil(0, val)
}

// IterAfter returns an unbuffered channel that sorted records can be read from and processed.
// IterAfter starts at the given value and sends all values until reaching the end of the collection.
func (sm *SortedMap) IterAfter(val interface{}) <-chan Record {
	return sm.iterAfter(0, val)
}

// BufferedIter returns a buffered channel that sorted records can be read from and processed.
func (sm *SortedMap) BufferedIter(bufSize int) <-chan Record {
	return sm.iter(bufSize)
}

// BufferedIterUntil returns a buffered channel that sorted records can be read from and processed.
// BufferedIterUntil starts at the lowest value in the collection and sends all values until reaching the given value.
func (sm *SortedMap) BufferedIterUntil(bufSize int, val interface{}) <-chan Record {
	return sm.iterUntil(bufSize, val)
}

// BufferedIterAfter returns a buffered channel that sorted records can be read from and processed.
// BufferedIterAfter starts at the given value and sends all values until reaching the end of the collection.
func (sm *SortedMap) BufferedIterAfter(bufSize int, val interface{}) <-chan Record {
	return sm.iterAfter(bufSize, val)
}