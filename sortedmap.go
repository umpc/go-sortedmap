package sortedmap

// SortedMap contains a map, slice, and reference to a sorting function.
// SortedMap is not concurrency-safe, though it can be easily wrapped by a developer-defined type.
type SortedMap struct {
	idx    map[string]interface{}
	sorted []string
	lessFn SortLessFunc
}

// Record defines a type used in batching and iterations, where keys and values are used together.
type Record struct {
	Key string
	Val interface{}
}

// SortLessFunc defines the type of the 'less than' comparison function for the default or chosen type.
type SortLessFunc func(idx map[string]interface{}, sorted []string, i int, val interface{}) bool

// New creates and initializes a new SortedMap structure and returns a reference to it.
func New(lessFn SortLessFunc) *SortedMap {
	lessFn = setDefaults(lessFn)

	return &SortedMap{
		idx:    make(map[string]interface{}),
		sorted: make([]string, 0),
		lessFn: lessFn,
	}
}

// Len returns the number of items in the collection.
func (sm *SortedMap) Len() int {
	return len(sm.sorted)
}