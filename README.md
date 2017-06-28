# SortedMap

[![Build Status](https://travis-ci.org/umpc/go-sortedmap.svg?branch=master)](https://travis-ci.org/umpc/go-sortedmap) [![Coverage Status](https://codecov.io/github/umpc/go-sortedmap/badge.svg?branch=master)](https://codecov.io/github/umpc/go-sortedmap?branch=master) [![GoDoc](https://godoc.org/github.com/umpc/go-sortedmap?status.svg)](https://godoc.org/github.com/umpc/go-sortedmap)

SortedMap is a small library that provides a value-sorted ```map[interface{}]interface{}``` type and methods combined from Go 1 map and slice primitives.

This data structure allows for roughly constant-time reads and for efficiently iterating over only a subsection of the stored values.

### Worst-Case Complexity
Operation | Worst-Case
----------|-----------
Has | ```O(1)```
Get | ```O(1)```
Iter | ```O(n)```
Delete | ```O(n log n)```
Insert | ```O(n^2)```
Replace | ```O(2^n)```

## Example Usage

```go
package main

import (
  "fmt"
  "time"
  mrand "math/rand"

  "github.com/umpc/go-sortedmap"
  "github.com/umpc/go-sortedmap/asc"
)

func main() {
  records := randRecords(23)

  // Create a new collection and set its insertion sort order:
  sm := sortedmap.New(asc.Time)
  sm.BatchInsert(records)

  // Loop through records, in order, until reaching the given value:
  for rec := range sm.IterUntil(time.Now()) {
    fmt.Printf("%+v\n", rec)
  }

  // Check out the docs for more functionality and further explainations.
}

func randRecords(n int) []*sortedmap.Record {
  mrand.Seed(time.Now().UTC().UnixNano())
  records := make([]*sortedmap.Record, n)
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
    min  := mrand.Intn(59)
    sec  := mrand.Intn(59)

    t := time.Date(year, mth, day, hour, min, sec, 0, time.UTC)

    records[i] = &sortedmap.Record{
      Key: t.Format(time.UnixDate),
      Val: t,
    }
  }
  return records
}
```

## Benchmarks

```sh
BenchmarkDelete1of1Record-8           	 5000000	       348 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of10Records-8         	 3000000	       593 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of100Records-8        	 2000000	       960 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of1000Records-8       	 1000000	      2255 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of10000Records-8      	  500000	      4346 ns/op	       0 B/op	       0 allocs/op

BenchmarkGet1Record-8                 	50000000	        34.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet10Records-8               	50000000	        32.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet100Records-8              	30000000	        35.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet1000Records-8             	50000000	        31.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet10000Records-8            	50000000	        31.7 ns/op	       0 B/op	       0 allocs/op

BenchmarkHas1Record-8                 	50000000	        33.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas10Records-8               	50000000	        32.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas100Records-8              	50000000	        35.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas1000Records-8             	50000000	        31.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas10000Records-8            	50000000	        34.4 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchInsert1Record-8         	30000000	        55.6 ns/op	       1 B/op	       1 allocs/op
BenchmarkBatchInsert10Records-8       	 3000000	       406 ns/op	      16 B/op	       1 allocs/op
BenchmarkBatchInsert100Records-8      	  300000	      4950 ns/op	     112 B/op	       1 allocs/op
BenchmarkBatchInsert1000Records-8     	   20000	     60208 ns/op	    1033 B/op	       1 allocs/op
BenchmarkBatchInsert10000Records-8    	    2000	    797372 ns/op	   11295 B/op	       1 allocs/op

BenchmarkBatchReplace1Record-8        	10000000	       250 ns/op	       0 B/op	       0 allocs/op
BenchmarkBatchReplace10Records-8      	  300000	      5784 ns/op	       0 B/op	       0 allocs/op
BenchmarkBatchReplace100Records-8     	   10000	    109066 ns/op	       0 B/op	       0 allocs/op
BenchmarkBatchReplace1000Records-8    	    1000	   1911582 ns/op	       0 B/op	       0 allocs/op
BenchmarkBatchReplace10000Records-8   	      20	  60068331 ns/op	       0 B/op	       0 allocs/op

BenchmarkNew-8                        	20000000	       119 ns/op	      96 B/op	       2 allocs/op
```

## License

The source code is available under the [MIT License](https://opensource.org/licenses/MIT).
