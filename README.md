# SortedMap

[![GoDoc](https://godoc.org/github.com/umpc/go-sortedmap?status.svg)](https://godoc.org/github.com/umpc/go-sortedmap)

Sorted Map is a small library that provides a value-sorted ```map[string]interface``` type and methods combined from Go 1 map and slice primitives.

This data structure allows for roughly constant-time reads and for efficiently iterating over only a subsection of the stored values. While this works for the project it was built for, because of the structure's reliance on a sorted slice of keys, worst-case delete operations are roughly ```O(n)```, where ```n``` is the number of items in the collection.

## Example Usage

```go
package main

import (
	"fmt"
	"time"
	mrand "math/rand"

	"github.com/umpc/go-sortedmap"
)

func main() {
	mrand.Seed(time.Now().UTC().UnixNano())

	// Example records:
	records := make([]*sortedmap.Record, 5)
	for i := range records {
		year := mrand.Intn(2058)
		for year < 2000 {
			year = mrand.Intn(2058)
		}
		mth := time.Month(mrand.Intn(12))
		if mth < 1 {
			mth++
		}
		day := mrand.Intn(28)
		if day < 1 {
			day++
		}

		hour := mrand.Intn(23)
		min := mrand.Intn(59)
		sec := mrand.Intn(59)
	
		t := time.Date(year, mth, day, hour, min, sec, 0, time.UTC)
		records[i] = &sortedmap.Record{
			Key: t.Format(time.RFC3339),
// 			Val: mrand.Intn(34334534561),
			Val: t,
		}
	}

	// Only one type can be used at a time, though handling for multiple types is still shown here:
	sm := sortedmap.New(func(idx map[string]interface{}, sorted []string, i int, val interface{}) bool {
		switch val.(type) {
//		case int:
//			return val.(int) < idx[sorted[i]].(int)
		case time.Time:
			return val.(time.Time).Before(idx[sorted[i]].(time.Time))
		default:
			return false
		}
	})

	// This instance is similar to the above declaration and contains the time.Time 'less than' conditional function shown above:
	// sm := sortedmap.New(nil)

	// Insert:
	sm.BatchReplace(records...)

	// Ordered iteration up until a given time:
	for rec := range sm.IterUntil(time.Now()) {
		fmt.Printf("%+v\n", rec)
	}
}
```

## License

The source code is available under the [MIT License](https://opensource.org/licenses/MIT).
