# SortedMap

[![Build Status](https://travis-ci.org/umpc/go-sortedmap.svg?branch=master)](https://travis-ci.org/umpc/go-sortedmap) [![Coverage Status](https://codecov.io/github/umpc/go-sortedmap/badge.svg?branch=master)](https://codecov.io/github/umpc/go-sortedmap?branch=master) [![GoDoc](https://godoc.org/github.com/umpc/go-sortedmap?status.svg)](https://godoc.org/github.com/umpc/go-sortedmap)

SortedMap is a small library that provides a value-sorted ```map[interface{}]interface{}``` type and methods combined from Go 1 map and slice primitives.

This data structure allows for roughly constant-time reads and for efficiently iterating over only a section of stored values.

### Complexity
Operation | Average-Case
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
BenchmarkDelete1of1Records-8                 	 5000000	       380 ns/op	       0 B/op	       0 allocs/op

BenchmarkDelete1of10Records-8                	 2000000	       834 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of100Records-8               	 1000000	      1443 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of1000Records-8              	  500000	      2400 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of10000Records-8             	  200000	      7318 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchDelete10of10Records-8          	  300000	      4198 ns/op	      16 B/op	       1 allocs/op
BenchmarkBatchDelete100of100Records-8        	   30000	     56898 ns/op	     112 B/op	       1 allocs/op
BenchmarkBatchDelete1000of1000Records-8      	    2000	    860714 ns/op	    1024 B/op	       1 allocs/op
BenchmarkBatchDelete10000of10000Records-8    	      50	  20510810 ns/op	   10240 B/op	       1 allocs/op

BenchmarkGet1of1Records-8                    	10000000	       122 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchGet10of10Records-8             	 1000000	      1177 ns/op	     176 B/op	       2 allocs/op
BenchmarkBatchGet100of100Records-8           	  200000	      7037 ns/op	    1904 B/op	       2 allocs/op
BenchmarkBatchGet1000of1000Records-8         	   30000	     59494 ns/op	   17408 B/op	       2 allocs/op
BenchmarkBatchGet10000of10000Records-8       	    2000	    848944 ns/op	  174080 B/op	       2 allocs/op

BenchmarkHas1of1Records-8                    	10000000	       123 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchHas10of10Records-8             	 2000000	       923 ns/op	      16 B/op	       1 allocs/op
BenchmarkBatchHas100of100Records-8           	  200000	      6347 ns/op	     112 B/op	       1 allocs/op
BenchmarkBatchHas1000of1000Records-8         	   30000	     54532 ns/op	    1024 B/op	       1 allocs/op
BenchmarkBatchHas10000of10000Records-8       	    2000	    630893 ns/op	   10240 B/op	       1 allocs/op

BenchmarkInsert1Record-8                     	 2000000	       598 ns/op	     304 B/op	       2 allocs/op

BenchmarkBatchInsert10Records-8              	  300000	      5030 ns/op	    1382 B/op	       8 allocs/op
BenchmarkBatchInsert100Records-8             	   20000	     64203 ns/op	   14908 B/op	      19 allocs/op
BenchmarkBatchInsert1000Records-8            	    2000	   1016292 ns/op	  202005 B/op	      78 allocs/op
BenchmarkBatchInsert10000Records-8           	      50	  23785169 ns/op	 2120592 B/op	     580 allocs/op

BenchmarkReplace1of1Records-8                	 3000000	       514 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplace1of10Records-8               	 1000000	      1322 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplace1of100Records-8              	 1000000	      2313 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplace1of1000Records-8             	  500000	      3880 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplace1of10000Records-8            	  200000	     10817 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchReplace10of10Records-8         	  300000	      4212 ns/op	      16 B/op	       1 allocs/op
BenchmarkBatchReplace100of100Records-8       	   20000	     59998 ns/op	     112 B/op	       1 allocs/op
BenchmarkBatchReplace1000of1000Records-8     	    2000	    858613 ns/op	    1024 B/op	       1 allocs/op
BenchmarkBatchReplace10000of10000Records-8   	     100	  20771583 ns/op	   10240 B/op	       1 allocs/op

BenchmarkNew-8                               	10000000	       124 ns/op	      96 B/op	       2 allocs/op
```

Better performance than the displayed benchmark test results is possible. The benchmark tests include some overhead from benchmark test functions that have been abstracted for easier comprehension and maintenance.

The above benchmark tests were ran on a 2.6GHz Intel Core i7-6700HQ (Skylake) CPU.

## License

The source code is available under the [MIT License](https://opensource.org/licenses/MIT).
