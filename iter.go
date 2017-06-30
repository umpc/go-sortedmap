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

// IterCallbackFunc defines the type of function that is passed into an IterFunc method and its single argument, a reference to a record value.
type IterCallbackFunc func(rec *Record) bool

func setBufSize(bufSize int) int {
	// initialBufSize must be >= 1 or a blocked channel send goroutine may not exit.
	// More info: https://github.com/golang/go/wiki/Timeouts
	const initialBufSize = 1

	if bufSize < initialBufSize {
		return initialBufSize
	}
	return bufSize
}

func (sm *SortedMap) sendRecord(ch chan Record, sendTimeout time.Duration, i int) bool {
	rec := Record{}

	rec.Key = sm.sorted[i]
	rec.Val = sm.idx[rec.Key]

	if sendTimeout == time.Duration(0) {
		ch <- rec
		return true
	}

	select {
	case ch <- rec:
		return true
	case <-time.After(sendTimeout):
		break
	}

	return false
}

func (sm *SortedMap) returnRecord(i int) *Record {
	rec := new(Record)
	rec.Key = sm.sorted[i]
	rec.Val = sm.idx[rec.Key]

	return rec
}

// parseChBoundParams returns a new reference to default parameters
// if no parameters reference was passed into the method.
func (sm *SortedMap) parseChBoundParams(params *IterChParams) *IterChParams {
	localParams := new(IterChParams)

	if params != nil {
		localParams.Reversed = params.Reversed
		localParams.SendTimeout = params.SendTimeout
		localParams.BufSize = params.BufSize

		iterBounds := sm.boundsIdxSearch(params.LowerBound, params.UpperBound)
		if len(iterBounds) >= 2 {
			localParams.LowerBound = iterBounds[0]
			localParams.UpperBound = iterBounds[1]
		}
	}

	return localParams
}

// (*IterChParams).bounds returns a slice containing the lower and upper bound indexes
// after being obtained and set using (*SortedMap).boundsIdxSearch
func (params *IterChParams) bounds() []int {
	switch params.LowerBound.(type) {
	case int:
	default:
		return nil
	}
	switch params.UpperBound.(type) {
	case int:
	default:
		return nil
	}
	return []int{
		params.LowerBound.(int),
		params.UpperBound.(int),
	}
}

func (sm *SortedMap) iterCh(params *IterChParams) <-chan Record {

	localParams := sm.parseChBoundParams(params)
	ch := make(chan Record, setBufSize(localParams.BufSize))

	go func(params *IterChParams, ch chan Record) {
		if params.LowerBound != nil && params.UpperBound != nil {
			iterBounds := params.bounds()

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
		} else if params.Reversed {
			for i := len(sm.sorted) - 1; i > 0; i-- {
				if !sm.sendRecord(ch, params.SendTimeout, i) {
					break
				}
			}
		} else {
			for i := range sm.sorted {
				if !sm.sendRecord(ch, params.SendTimeout, i) {
					break
				}
			}
		}
		close(ch)
	}(localParams, ch)

	return ch
}

func (sm *SortedMap) iterFunc(reversed bool, lowerBound, upperBound interface{}, f func(rec *Record) bool) {
	iterBounds := sm.boundsIdxSearch(lowerBound, upperBound)
	if len(iterBounds) < 2 {
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
	} else if reversed {
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
// This method defaults to the expected behavior of blocking until a read, with no timeout.
func (sm *SortedMap) IterCh() <-chan Record {
	return sm.iterCh(nil)
}

// BoundedIterCh returns a channel that sorted records can be read from and processed.
// BoundedIterCh starts at the lower bound value and sends all values in the collection until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
// This method defaults to the expected behavior of blocking until a channel send completes, with no timeout.
func (sm *SortedMap) BoundedIterCh(reversed bool, lowerBound, upperBound interface{}) <-chan Record {
	return sm.iterCh(&IterChParams{
		Reversed: reversed,
		LowerBound: lowerBound,
		UpperBound: upperBound,
	})
}

// CustomIterCh returns a channel that sorted records can be read from and processed.
// CustomIterCh starts at the lower bound value and sends all values in the collection until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
// This method defaults to the expected behavior of blocking until a channel send completes, with no timeout.
func (sm *SortedMap) CustomIterCh(params *IterChParams) <-chan Record {
	return sm.iterCh(params)
}

// IterFunc passes each record to the specified callback function.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap) IterFunc(reversed bool, f IterCallbackFunc) {
	sm.iterFunc(reversed, nil, nil, f)
}

// BoundedIterFunc starts at the lower bound value and passes all values in the collection to the callback function until reaching the upper bounds value.
// Sort order is reversed if the reversed argument is set to true.
func (sm *SortedMap) BoundedIterFunc(reversed bool, lowerBound, upperBound interface{}, f IterCallbackFunc) {
	sm.iterFunc(reversed, lowerBound, upperBound, f)
}