package sortedmap

import (
	"errors"
	"time"
)

// IterChCloser allows records to be read through a channel that is returned by the Records method.
// IterChCloser values should be closed after use using the Close method.
type IterChCloser struct {
	ch       chan Record
	canceled chan struct{}
}

// Close cancels a channel-based iteration and causes the sending goroutine to exit.
// Close should be used after an IterChCloser is finished being read from.
func (iterCh *IterChCloser) Close() error {
	select {
	case iterCh.canceled <- struct{}{}:
	default:
	}

	return nil
}

// Records returns nil if the IterChCloser has been closed.
// Otherwise, Record returns a channel that records can be read from.
func (iterCh *IterChCloser) Records() <-chan Record {
	select {
	case <-iterCh.canceled:
		iterCh.canceled <- struct{}{}
		return nil
	default:
		return iterCh.ch
	}
}

// IterChParams contains configurable settings for CustomIterCh.
// SendTimeout is disabled by default, though it should be set to allow
// channel send goroutines to time-out.
// BufSize is set to 1 if its field is set to a lower value.
// LowerBound and UpperBound default to regular iteration when left unset.
type IterChParams struct {
	Reversed    bool
	SendTimeout time.Duration
	BufSize     int
	LowerBound,
	UpperBound interface{}
}

// IterCallbackFunc defines the type of function that is passed into an IterFunc method.
// The function is passed a record value argument.
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

func (sm *SortedMap) sendRecord(iterCh IterChCloser, sendTimeout time.Duration, i int) bool {
	select {
	case <-iterCh.canceled:
		iterCh.canceled <- struct{}{}
		return false
	default:
	}

	if sendTimeout <= time.Duration(0) {
		iterCh.ch <- sm.recordFromIdx(i)
		return true
	}

	select {
	case iterCh.ch <- sm.recordFromIdx(i):
		return true

	case <-time.After(sendTimeout):
		return false
	}
}

func (sm *SortedMap) iterCh(params IterChParams) (IterChCloser, error) {

	iterBounds := sm.boundsIdxSearch(params.LowerBound, params.UpperBound)
	if iterBounds == nil {
		return IterChCloser{}, errors.New(noValuesErr)
	}

	iterCh := IterChCloser{
		ch:       make(chan Record, setBufSize(params.BufSize)),
		canceled: make(chan struct{}, 1),
	}

	go func(params IterChParams, iterCh IterChCloser) {
		if params.Reversed {
			for i := iterBounds[1]; i >= iterBounds[0]; i-- {
				if !sm.sendRecord(iterCh, params.SendTimeout, i) {
					break
				}
			}
		} else {
			for i := iterBounds[0]; i <= iterBounds[1]; i++ {
				if !sm.sendRecord(iterCh, params.SendTimeout, i) {
					break
				}
			}
		}
		close(iterCh.ch)
	}(params, iterCh)

	return iterCh, nil
}

func (sm *SortedMap) iterFunc(reversed bool, lowerBound, upperBound interface{}, f IterCallbackFunc) error {

	iterBounds := sm.boundsIdxSearch(lowerBound, upperBound)
	if iterBounds == nil {
		return errors.New(noValuesErr)
	}

	if reversed {
		for i := iterBounds[1]; i >= iterBounds[0]; i-- {
			if !f(sm.recordFromIdx(i)) {
				break
			}
		}
	} else {
		for i := iterBounds[0]; i <= iterBounds[1]; i++ {
			if !f(sm.recordFromIdx(i)) {
				break
			}
		}
	}

	return nil
}

// IterCh returns a channel that sorted records can be read from and processed.
// This method defaults to the expected behavior of blocking until a read, with no timeout.
func (sm *SortedMap) IterCh() (IterChCloser, error) {
	return sm.iterCh(IterChParams{})
}

// BoundedIterCh returns a channel that sorted records can be read from and processed.
// BoundedIterCh starts at the lower bound value and sends all values in the collection until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
// This method defaults to the expected behavior of blocking until a channel send completes, with no timeout.
func (sm *SortedMap) BoundedIterCh(reversed bool, lowerBound, upperBound interface{}) (IterChCloser, error) {
	return sm.iterCh(IterChParams{
		Reversed:   reversed,
		LowerBound: lowerBound,
		UpperBound: upperBound,
	})
}

// CustomIterCh returns a channel that sorted records can be read from and processed.
// CustomIterCh starts at the lower bound value and sends all values in the collection until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
// This method defaults to the expected behavior of blocking until a channel send completes, with no timeout.
func (sm *SortedMap) CustomIterCh(params IterChParams) (IterChCloser, error) {
	return sm.iterCh(params)
}

// IterFunc passes each record to the specified callback function.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap) IterFunc(reversed bool, f IterCallbackFunc) {
	sm.iterFunc(reversed, nil, nil, f)
}

// BoundedIterFunc starts at the lower bound value and passes all values in the collection to the callback function until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap) BoundedIterFunc(reversed bool, lowerBound, upperBound interface{}, f IterCallbackFunc) error {
	return sm.iterFunc(reversed, lowerBound, upperBound, f)
}
