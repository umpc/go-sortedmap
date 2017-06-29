package sortedmap

func (sm *SortedMap) sendRecord(ch chan Record, i int) {
	rec := Record{
		Key: sm.sorted[i],
	}
	rec.Val = sm.idx[rec.Key]
	ch <- rec
}

func (sm *SortedMap) returnRecord(i int) Record {
	rec := Record{
		Key: sm.sorted[i],
	}
	rec.Val = sm.idx[rec.Key]
	return rec
}

func (sm *SortedMap) iterCh(reversed bool, bufSize int) <-chan Record {
	ch := make(chan Record, bufSize)

	go func(ch chan Record) {
		if reversed {
			for i := len(sm.sorted) - 1; i > 0; i-- {
				sm.sendRecord(ch, i)
			}
		} else {
			for i := range sm.sorted {
				sm.sendRecord(ch, i)
			}
		}
		close(ch)
	}(ch)

	return ch
}

func (sm *SortedMap) iterBetweenCh(reversed bool, bufSize int, lowerBound, upperBound interface{}) (<-chan Record, bool) {

	iterBounds := sm.between(lowerBound, upperBound)
	if len(iterBounds) < 2 {
		return nil, false
	}

	ch := make(chan Record, bufSize)
	go func(reversed bool, iterBounds []int, ch chan Record) {
		if reversed {
			for i := iterBounds[1]; i > iterBounds[0]; i-- {
				sm.sendRecord(ch, i)
			}
		} else {
			for i := iterBounds[0]; i < iterBounds[1]; i++ {
				sm.sendRecord(ch, i)
			}
		}
		close(ch)
	}(reversed, iterBounds, ch)

	return ch, true
}

func (sm *SortedMap) iterFunc(reversed bool, f func(rec Record) bool) {
	if reversed {
		for i := len(sm.sorted) - 1; i > 0; i-- {
			if !f(sm.returnRecord(i)) {
				break
			}
		}
	} else {
		for i := range sm.sorted {
			if !f(sm.returnRecord(i)) {
				break
			}
		}
	}
}

func (sm *SortedMap) iterBetweenFunc(reversed bool, lowerBound, upperBound interface{}, f func(rec Record) bool) {
	iterBounds := sm.between(lowerBound, upperBound)
	if len(iterBounds) < 2 {
		return
	}
	if reversed {
		for i := iterBounds[1]; i > iterBounds[0]; i-- {
			if !f(sm.returnRecord(i)) {
				break
			}
		}
	} else {
		for i := iterBounds[0]; i < iterBounds[1]; i++ {
			if !f(sm.returnRecord(i)) {
				break
			}
		}
	}
}

// IterCh returns a channel that sorted records can be read from and processed.
func (sm *SortedMap) IterCh() <-chan Record {
	return sm.iterCh(false, 0)
}

// IterChCustom returns a channel that sorted records can be read from and processed, with custom options.
func (sm *SortedMap) IterChCustom(reversed bool, bufSize int) <-chan Record {
	return sm.iterCh(reversed, bufSize)
}

// IterBetweenCh returns a channel that sorted records can be read from and processed,
// and a boolean value that indicates whether or not values in the collection fall between the given bounds.
// IterBetweenCh starts at the lower bound value and sends all values in the collection until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap) IterBetweenCh(lowerBound, upperBound interface{}) (<-chan Record, bool) {
	return sm.iterBetweenCh(false, 0, lowerBound, upperBound)
}

// IterBetweenChCustom returns a channel that sorted records can be read from and processed,
// and a boolean value that indicates whether or not values in the collection fall between the given bounds.
// IterBetweenChCustom starts at the lower bound value and sends all values in the collection until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap) IterBetweenChCustom(reversed bool, bufSize int, lowerBound, upperBound interface{}) (<-chan Record, bool) {
	return sm.iterBetweenCh(reversed, bufSize, lowerBound, upperBound)
}

// IterFunc passes each record to the specified callback function.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap) IterFunc(reversed bool, f func(rec Record) bool) {
	sm.iterFunc(reversed, f)
}

// IterBetweenFunc starts at the lower bound value and passes all values in the collection to the callback function until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap) IterBetweenFunc(reversed bool, lowerBound, upperBound interface{}, f func(rec Record) bool) {
	sm.iterBetweenFunc(reversed, lowerBound, upperBound, f)
}