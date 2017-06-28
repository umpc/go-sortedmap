package sortedmap

import "sort"

func (sm *SortedMap) iter(bufSize int) <-chan Record {
	ch := make(chan Record, bufSize)

	go func(ch chan Record) {
		smLen := len(sm.sorted)
		for i := 0; i < smLen; i++ {
			rec := Record{
				Key: sm.sorted[i],
			}
			rec.Val = sm.idx[rec.Key]

			ch <- rec
		}
		close(ch)
	}(ch)

	return ch
}

func (sm *SortedMap) iterUntil(bufSize int, val interface{}) <-chan Record {
	ch := make(chan Record, bufSize)

	go func(ch chan Record) {
		smLen := len(sm.sorted)
		for i := 0; i < smLen; i++ {
			rec := Record{
				Key: sm.sorted[i],
			}
			rec.Val = sm.idx[rec.Key]

			if sm.lessFn(val, rec.Val) {
				break
			}
			ch <- rec
		}
		close(ch)
	}(ch)

	return ch
}

func (sm *SortedMap) iterAfter(bufSize int, val interface{}) <-chan Record {
	ch := make(chan Record, bufSize)

	go func(ch chan Record) {
		smLen := len(sm.sorted)
		i := sort.Search(smLen, func(i int) bool {
			return sm.lessFn(val, sm.idx[sm.sorted[i]])
		})
		for; i < smLen; i++ {
			rec := Record{
				Key: sm.sorted[i],
			}
			rec.Val = sm.idx[rec.Key]

			ch <- rec
		}
		close(ch)
	}(ch)

	return ch
}

func (sm *SortedMap) reverseIter(bufSize int) <-chan Record {
	ch := make(chan Record, bufSize)

	go func(ch chan Record) {
		smLen := len(sm.sorted)
		for i := smLen - 1; i > 0; i-- {
			rec := Record{
				Key: sm.sorted[i],
			}
			rec.Val = sm.idx[rec.Key]

			ch <- rec
		}
		close(ch)
	}(ch)

	return ch
}

func (sm *SortedMap) reverseIterUntil(bufSize int, val interface{}) <-chan Record {
	ch := make(chan Record, bufSize)

	go func(ch chan Record) {
		smLen := len(sm.sorted)
		for i := smLen - 1; i > 0; i-- {
			rec := Record{
				Key: sm.sorted[i],
			}
			rec.Val = sm.idx[rec.Key]

			if sm.lessFn(val, rec.Val) {
				break
			}
			ch <- rec
		}
		close(ch)
	}(ch)

	return ch
}

func (sm *SortedMap) reverseIterAfter(bufSize int, val interface{}) <-chan Record {
	ch := make(chan Record, bufSize)

	go func(ch chan Record) {
		smLen := len(sm.sorted)
		i := sort.Search(smLen, func(i int) bool {
			return sm.lessFn(val, sm.idx[sm.sorted[i]])
		})
		for; i > 0; i-- {
			rec := Record{
				Key: sm.sorted[i],
			}
			rec.Val = sm.idx[rec.Key]

			ch <- rec
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

// ReverseIter returns an unbuffered channel that reverse-sorted records can be read from and processed.
func (sm *SortedMap) ReverseIter() <-chan Record {
	return sm.reverseIter(0)
}

// ReverseIterUntil returns an unbuffered channel that reverse-sorted records can be read from and processed.
// ReverseIterUntil starts at the highest value in the collection and sends all values until reaching the given value.
func (sm *SortedMap) ReverseIterUntil(val interface{}) <-chan Record {
	return sm.reverseIterUntil(0, val)
}

// ReverseIterAfter returns an unbuffered channel that reverse-sorted records can be read from and processed.
// ReverseIterAfter starts at the given value and sends all values until reaching the beginning of the collection.
func (sm *SortedMap) ReverseIterAfter(val interface{}) <-chan Record {
	return sm.reverseIterAfter(0, val)
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

// BufferedReverseIter returns a buffered channel that reverse-sorted records can be read from and processed.
func (sm *SortedMap) BufferedReverseIter(bufSize int) <-chan Record {
	return sm.reverseIter(bufSize)
}

// BufferedReverseIterUntil returns a buffered channel that reverse-sorted records can be read from and processed.
// BufferedReverseIterUntil starts at the highest value in the collection and sends all values until reaching the given value.
func (sm *SortedMap) BufferedReverseIterUntil(bufSize int, val interface{}) <-chan Record {
	return sm.reverseIterUntil(bufSize, val)
}

// BufferedReverseIterAfter returns a buffered channel that reverse-sorted records can be read from and processed.
// BufferedReverseIterAfter starts at the given value and sends all values until reaching the beginning of the collection.
func (sm *SortedMap) BufferedReverseIterAfter(bufSize int, val interface{}) <-chan Record {
	return sm.reverseIterAfter(bufSize, val)
}