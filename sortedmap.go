package sortedmap

type SortedMap struct {
	idx    map[string]interface{}
	sorted []string
	lessFn SortLessFunc
}

type Record struct {
	Key string
	Val interface{}
}

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