package sortedmap

import "sort"

func (sm *SortedMap) between(lowerBound, upperBound interface{}) []int {

	if lowerBound == nil || upperBound == nil {
		return nil
	}

	smLen := len(sm.sorted)
	lowerBoundIdx := sort.Search(smLen, func(i int) bool {
		return sm.lessFn(lowerBound, sm.idx[sm.sorted[i]])
	})

	// Return if lowerBoundIdx is higher than all other slice indexes.
	if lowerBoundIdx == smLen {
		return nil
	}

	upperBoundIdx := sort.Search(smLen, func(i int) bool {
		return sm.lessFn(upperBound, sm.idx[sm.sorted[i]])
	})

	// sort.Search returns the smallest index i in [0, n) at which f(i) is true.
	// This check stops the deletion of values larger than the upperBound value.
	if upperBoundIdx == smLen {
		upperBoundIdx -= 2
	} else if upperBoundIdx > 0 {
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