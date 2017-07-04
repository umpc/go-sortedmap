package sortedmap

import "sort"

func (sm *SortedMap) setBoundIdx(boundVal interface{}) int {

	if boundVal == nil {
		return 0
	}

	smLen := len(sm.sorted)
	idx := sort.Search(smLen, func(i int) bool {
		return sm.lessFn(boundVal, sm.idx[sm.sorted[i]])
	})

	// sort.Search returns the smallest index i in [0, n) at which f(i) is true.
	// This sets the correct index for less than conditional comparisons.
	if idx > 0 {
		idx--
	}

	return idx
}

func (sm *SortedMap) boundsIdxSearch(lowerBound, upperBound interface{}) []int {
	smLen := len(sm.sorted)

	if lowerBound != nil && upperBound != nil {
		if sm.lessFn(upperBound, lowerBound) {
			return nil
		}
	}

	lowerBoundIdx := sm.setBoundIdx(lowerBound)
	if lowerBound != nil {
		if lowerBoundIdx < smLen-1 {
			valFromIdx := sm.idx[sm.sorted[lowerBoundIdx]]

			// If the bound value is greater than or equal to the value from the map,
			// select the next index value.
			if sm.lessFn(valFromIdx, lowerBound) {
				lowerBoundIdx++
			}
		}
	}

	upperBoundIdx := 0
	if upperBound == nil {
		upperBoundIdx = smLen - 1
	} else {
		upperBoundIdx = sm.setBoundIdx(upperBound)
		if upperBoundIdx < smLen-1 {
			valFromIdx := sm.idx[sm.sorted[upperBoundIdx + 1]]
			if !sm.lessFn(valFromIdx, upperBound) && !sm.lessFn(upperBound, valFromIdx) {
				upperBoundIdx++
			}
		}
	}

	if lowerBound != nil && upperBound != nil {
		if lowerBoundIdx == upperBoundIdx {
			valFromIdx := sm.idx[sm.sorted[lowerBoundIdx]]

			if !sm.lessFn(lowerBound, upperBound) && !sm.lessFn(upperBound, lowerBound) {
				// lowerBound == upperBound
				if sm.lessFn(valFromIdx, lowerBound) || sm.lessFn(lowerBound, valFromIdx) {
					return nil
				}
			}

			if sm.lessFn(valFromIdx, lowerBound) || sm.lessFn(upperBound, valFromIdx) {
				return nil
			}
		}
	}

	return []int{
		lowerBoundIdx,
		upperBoundIdx,
	}
}
