package sortedmap

import "time"

// IterChParams contains configurable settings for CustomIterCh.
// SendTimeout is disabled by default, though it should be set to allow
// channel send goroutines to time-out.
// BufSize is set to 1 if its field is set to a lower value.
// LowerBound and UpperBound default to regular iteration when left unset.
type IterChParams struct{
	Reversed bool
	SendTimeout time.Duration
	BufSize int
	LowerBound,
	UpperBound interface{}
}

// IterCallbackFunc defines the type of function that is passed into an IterFunc method and its single argument, which is a reference to a record value.
type IterCallbackFunc func(rec Record) bool

func setBufSize(bufSize int) int {
	// initialBufSize must be >= 1 or a blocked channel send goroutine may not exit.
	// More info: https://github.com/golang/go/wiki/Timeouts
	const initialBufSize = 1

	if bufSize < initialBufSize {
		return initialBufSize
	}
	return bufSize
}

func (sm *SortedMap) recordFromIdx(i int) Record {
	rec := Record{}
	rec.Key = sm.sorted[i]
	rec.Val = sm.idx[rec.Key]

	return rec
}

func (sm *SortedMap) sendRecord(ch chan Record, sendTimeout time.Duration, i int) bool {

	if sendTimeout <= time.Duration(0) {
		ch <- sm.recordFromIdx(i)
		return true
	}

	select {
	case ch <- sm.recordFromIdx(i):
		return true
	case <-time.After(sendTimeout):
		return false
	}
}

func (sm *SortedMap) iterCh(params IterChParams) (<-chan Record, bool) {

	iterBounds := sm.boundsIdxSearch(params.LowerBound, params.UpperBound)
	if len(iterBounds) < 2 {
		return nil, false
	}
	ch := make(chan Record, setBufSize(params.BufSize))

	go func(params IterChParams, ch chan Record) {
		if params.Reversed {
			for i := iterBounds[1]; i > iterBounds[0]; i-- {
				if !sm.sendRecord(ch, params.SendTimeout, i) {
					break
				}
			}
		} else {
			for i := iterBounds[0]; i < iterBounds[1]; i++ {
				if !sm.sendRecord(ch, params.SendTimeout, i) {
					break
				}
			}
		}
		close(ch)
	}(params, ch)

	return ch, true
}

func (sm *SortedMap) iterFunc(reversed bool, lowerBound, upperBound interface{}, f IterCallbackFunc) bool {

	iterBounds := sm.boundsIdxSearch(lowerBound, upperBound)
	if len(iterBounds) < 2 {
		return false
	}

	if reversed {
		for i := iterBounds[1]; i > iterBounds[0]; i-- {
			if !f(sm.recordFromIdx(i)) {
				break
			}
		}
	} else {
		for i := iterBounds[0]; i < iterBounds[1]; i++ {
			if !f(sm.recordFromIdx(i)) {
				break
			}
		}
	}
	return true
}

// IterCh returns a channel that sorted records can be read from and processed.
// This method defaults to the expected behavior of blocking until a read, with no timeout.
func (sm *SortedMap) IterCh() <-chan Record {
	ch, _ := sm.iterCh(IterChParams{})
	return ch
}

// BoundedIterCh returns a channel that sorted records can be read from and processed.
// BoundedIterCh starts at the lower bound value and sends all values in the collection until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
// This method defaults to the expected behavior of blocking until a channel send completes, with no timeout.
func (sm *SortedMap) BoundedIterCh(reversed bool, lowerBound, upperBound interface{}) (<-chan Record, bool) {
	return sm.iterCh(IterChParams{
		Reversed: reversed,
		LowerBound: lowerBound,
		UpperBound: upperBound,
	})
}

// CustomIterCh returns a channel that sorted records can be read from and processed.
// CustomIterCh starts at the lower bound value and sends all values in the collection until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
// This method defaults to the expected behavior of blocking until a channel send completes, with no timeout.
func (sm *SortedMap) CustomIterCh(params IterChParams) (<-chan Record, bool) {
	return sm.iterCh(params)
}

// IterFunc passes each record to the specified callback function.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap) IterFunc(reversed bool, f IterCallbackFunc) bool {
	return sm.iterFunc(reversed, nil, nil, f)
}

// BoundedIterFunc starts at the lower bound value and passes all values in the collection to the callback function until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap) BoundedIterFunc(reversed bool, lowerBound, upperBound interface{}, f IterCallbackFunc) bool {
	return sm.iterFunc(reversed, lowerBound, upperBound, f)
}