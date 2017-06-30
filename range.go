package sortedmap

import "sort"

func (sm *SortedMap) rangeIdxSearch(lowerBound, upperBound interface{}) []int {

	if lowerBound == nil || upperBound == nil {
		return nil
	}

	smLen := len(sm.sorted)
	lowerBoundIdx := sort.Search(smLen, func(i int) bool {
		return sm.lessFn(lowerBound, sm.idx[sm.sorted[i]])
	})

	// Lower bound is the largest value. Select the last index.
	if lowerBoundIdx == smLen {
		lowerBoundIdx--
	}

	upperBoundIdx := sort.Search(smLen, func(i int) bool {
		return sm.lessFn(upperBound, sm.idx[sm.sorted[i]])
	})

	if upperBoundIdx == smLen {
		upperBoundIdx--
	}

	// sort.Search returns the smallest index i in [0, n) at which f(i) is true.
	// This check stops the deletion of values larger than the upperBound value.
	if upperBoundIdx > 0 {
		upperBoundIdx--
	}

	if lowerBoundIdx > upperBoundIdx {
		return []int{
			upperBoundIdx,
			lowerBoundIdx,
		}
	} else {
		return []int{
			lowerBoundIdx,
			upperBoundIdx,
		}
	}
}