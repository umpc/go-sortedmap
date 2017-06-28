# SortedMap

[![Build Status](https://travis-ci.org/umpc/go-sortedmap.svg?branch=master)](https://travis-ci.org/umpc/go-sortedmap) [![Coverage Status](https://codecov.io/github/umpc/go-sortedmap/badge.svg?branch=master)](https://codecov.io/github/umpc/go-sortedmap?branch=master) [![GoDoc](https://godoc.org/github.com/umpc/go-sortedmap?status.svg)](https://godoc.org/github.com/umpc/go-sortedmap)

SortedMap is a small library that provides a value-sorted ```map[interface{}]interface{}``` type and methods combined from Go 1 map and slice primitives.

This data structure allows for roughly constant-time reads and for efficiently iterating over only a section of stored values.

### Complexity
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
BenchmarkDelete1of1Record-8                  	 5000000	       386 ns/o       0 B/op	       0 allocs/op
BenchmarkDelete1of10Records-8                	 3000000	       508 ns/o       0 B/op	       0 allocs/op
BenchmarkDelete1of100Records-8               	 2000000	       893 ns/o       0 B/op	       0 allocs/op
BenchmarkDelete1of1000Records-8              	 1000000	      1975 ns/o       0 B/op	       0 allocs/op
BenchmarkDelete1of10000Records-8             	  500000	      3168 ns/o       0 B/op	       0 allocs/op

BenchmarkGet1of1Record-8                     	50000000	        33.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet1of10Records-8                   	50000000	        29.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet1of100Records-8                  	50000000	        30.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet1of1000Records-8                 	50000000	        32.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet1of10000Records-8                	50000000	        33.3 ns/op	       0 B/op	       0 allocs/op

BenchmarkHas1of1Record-8                     	50000000	        34.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas1of10Records-8                   	50000000	        32.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas1of100Records-8                  	50000000	        32.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas1of1000Records-8                 	50000000	        31.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas1of10000Records-8                	50000000	        33.3 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchInsert1Record-8                	20000000	        56.1 ns/op	       1 B/op	       1 allocs/op
BenchmarkBatchInsert10Records-8              	 3000000	       397 ns/o      16 B/op	       1 allocs/op
BenchmarkBatchInsert100Records-8             	  300000	      4572 ns/o     112 B/op	       1 allocs/op
BenchmarkBatchInsert1000Records-8            	   30000	     55933 ns/o    1030 B/op	       1 allocs/op
BenchmarkBatchInsert10000Records-8           	    2000	    677099 ns/o   11293 B/op	       1 allocs/op

BenchmarkBatchReplace1of1Record-8            	10000000	       205 ns/o       0 B/op	       0 allocs/op
BenchmarkBatchReplace10of10Records-8         	  300000	      5457 ns/o       0 B/op	       0 allocs/op
BenchmarkBatchReplace100of100Records-8       	   10000	    107118 ns/o       0 B/op	       0 allocs/op
BenchmarkBatchReplace1000of1000Records-8     	    1000	   1982158 ns/o       0 B/op	       0 allocs/op
BenchmarkBatchReplace10000of10000Records-8   	      20	  68456103 ns/o       0 B/op	       0 allocs/op

BenchmarkNew-8                               	10000000	       119 ns/o      96 B/op	       2 allocs/op
```

The above benchmarks were ran on a 2.6GHz Intel Core i7-6700HQ (Skylake) CPU.

## License

The source code is available under the [MIT License](https://opensource.org/licenses/MIT).
