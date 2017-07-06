package sortedmap

// SortedMap contains a map, a slice, and references to one or more comparison functions.
// SortedMap is not concurrency-safe, though it can be easily wrapped by a developer-defined type.
type SortedMap struct {
	idx    map[interface{}]interface{}
	sorted []interface{}
	lessFn ComparisonFunc
}

// Record defines a type used in batching and iterations, where keys and values are used together.
type Record struct {
	Key,
	Val interface{}
}

// ComparisonFunc defines the type of the comparison function for the chosen value type.
type ComparisonFunc func(i, j interface{}) bool

func noOpComparisonFunc(_, _ interface{}) bool {
	return false
}

func setComparisonFunc(cmpFn ComparisonFunc) ComparisonFunc {
	if cmpFn == nil {
		return noOpComparisonFunc
	}
	return cmpFn
}

// New creates and initializes a new SortedMap structure and then returns a reference to it.
// New SortedMaps are created with a backing map/slice of length/capacity n.
func New(n int, cmpFn ComparisonFunc) *SortedMap {
	return &SortedMap{
		idx:    make(map[interface{}]interface{}, n),
		sorted: make([]interface{}, 0, n),
		lessFn: setComparisonFunc(cmpFn),
	}
}

// Len returns the number of items in the collection.
func (sm *SortedMap) Len() int {
	return len(sm.sorted)
}
