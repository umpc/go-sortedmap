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

func (sm *SortedMap) lowerBoundIncrCondition(i, j interface{}) bool {
	return sm.lessFn(j, i)
}

func (sm *SortedMap) lowerBoundRetCondition(i, j interface{}) bool {
	return sm.lessFn(j, i)
}

func (sm *SortedMap) upperBoundIncrCondition(i, j interface{}) bool {
	return !sm.lessFn(j, i) && !sm.lessFn(i, j)
}

func (sm *SortedMap) upperBoundRetCondition(i, j interface{}) bool {
	return sm.lessFn(i, j)
}

func (sm *SortedMap) adjustIdxOrReturn(boundVal interface{}, incrCondition, retCondition func(boundVal, valFromIdx interface{}) bool) (int, bool) {
	idx := sm.setBoundIdx(boundVal)
	if idx < len(sm.sorted)-1 {
		valFromIdx := sm.idx[sm.sorted[idx]]

		if incrCondition(boundVal, valFromIdx) {
			idx++
		}

		valFromIdx = sm.idx[sm.sorted[idx]]
		if retCondition(boundVal, valFromIdx) {
			return idx, false
		}
	}
	return idx, true
}

func (sm *SortedMap) boundsIdxSearch(lowerBound, upperBound interface{}) []int {
	smLen := len(sm.sorted)
	if smLen == 0 {
		return nil
	}

	if lowerBound != nil && upperBound != nil {
		if sm.lessFn(upperBound, lowerBound) {
			return nil
		}
	}

	lowerBoundIdx := sm.setBoundIdx(lowerBound)
	if lowerBound != nil {
		var ok bool
		lowerBoundIdx, ok = sm.adjustIdxOrReturn(lowerBound, sm.lowerBoundIncrCondition, sm.lowerBoundRetCondition)
		if !ok {
			return nil
		}
	}

	upperBoundIdx := 0
	if upperBound == nil {
		upperBoundIdx = smLen - 1
	} else {
		var ok bool
		upperBoundIdx, ok = sm.adjustIdxOrReturn(upperBound, sm.upperBoundIncrCondition, sm.upperBoundRetCondition)
		if !ok {
			return nil
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

	if lowerBoundIdx > upperBoundIdx {
		return nil
	}

	return []int{
		lowerBoundIdx,
		upperBoundIdx,
	}
}
