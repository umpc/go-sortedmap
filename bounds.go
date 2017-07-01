package sortedmap

import "sort"

func (sm *SortedMap) setBoundIdx(bound interface{}) int {
	idx := 0
	if bound == nil {
		return idx
	}
	smLen := len(sm.sorted)
	idx = sort.Search(smLen, func(i int) bool {
		return sm.lessFn(bound, sm.idx[sm.sorted[i]])
	})
	// sort.Search returns the smallest index i in [0, n) at which f(i) is true.
	// This sets the correct index for less than conditional comparisons.
	if idx > 0 {
		idx--
	}
	valFromIdx := sm.idx[sm.sorted[idx]]

	if idx < smLen - 1 && valFromIdx != bound {
		if !sm.lessFn(bound, valFromIdx) {
			idx++
		}
	}
	return idx
}

func (sm *SortedMap) boundsIdxSearch(lowerBound, upperBound interface{}) []int {
	lowerBoundIdx := sm.setBoundIdx(lowerBound)

	upperBoundIdx := 0
	if upperBound == nil {
		upperBoundIdx = len(sm.sorted) - 1
	} else {
		upperBoundIdx = sm.setBoundIdx(upperBound)
	}

<<<<<<< HEAD
	if lowerBound != nil && upperBound != nil {
		if lowerBoundIdx == upperBoundIdx {
			valFromIdx := sm.idx[sm.sorted[lowerBoundIdx]]

			if sm.lessFn(lowerBound, valFromIdx) {
				if sm.lessFn(upperBound, valFromIdx) {
					return nil
				}
=======
	if lowerBoundIdx == upperBoundIdx {
		valFromIdx := sm.idx[sm.sorted[lowerBoundIdx]]
		if sm.lessFn(lowerBound, valFromIdx) {
			if sm.lessFn(upperBound, valFromIdx) {
				return nil
>>>>>>> parent of 17c7ac5... Add missing nil check to the bounds index search method
			}
		}
		if !sm.lessFn(lowerBound, valFromIdx) {
			if !sm.lessFn(upperBound, valFromIdx) {
				return nil
			}
		}
	}

	return []int{
		lowerBoundIdx,
		upperBoundIdx,
	}
}