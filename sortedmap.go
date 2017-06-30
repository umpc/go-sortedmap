package sortedmap

// SortedMap contains a map, a slice, and a reference to a sorting function.
// SortedMap is not concurrency-safe, though it can be easily wrapped by a developer-defined type.
type SortedMap struct {
	idx    map[interface{}]interface{}
	sorted []interface{}
	lessFn SortLessFunc
}

// Record defines a type used in batching and iterations, where keys and values are used together.
type Record struct {
	Key,
	Val interface{}
}

// SortLessFunc defines the type of the comparison function for the chosen value type.
type SortLessFunc func(i, j interface{}) bool

func unsortedSortLessFunc(_, _ interface{}) bool {
	return false
}

func setLessFunc(lessFn SortLessFunc) SortLessFunc {
	if lessFn == nil {
		return unsortedSortLessFunc
	}
	return lessFn
}

// New creates and initializes a new SortedMap structure with a preallocated slice of capacity n, and returns a reference to it.
func New(n int, lessFn SortLessFunc) *SortedMap {
	return &SortedMap{
		idx:    make(map[interface{}]interface{}),
		sorted: make([]interface{}, 0, n),
		lessFn: setLessFunc(lessFn),
	}
}

// Len returns the number of items in the collection.
func (sm *SortedMap) Len() int {
	return len(sm.sorted)
}