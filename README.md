# SortedMap

[![Build Status](https://travis-ci.org/umpc/go-sortedmap.svg?branch=master)](https://travis-ci.org/umpc/go-sortedmap) [![Coverage Status](https://codecov.io/github/umpc/go-sortedmap/badge.svg?branch=master)](https://codecov.io/github/umpc/go-sortedmap?branch=master) [![GoDoc](https://godoc.org/github.com/umpc/go-sortedmap?status.svg)](https://godoc.org/github.com/umpc/go-sortedmap)

SortedMap is a small library that provides a value-sorted ```map[interface{}]interface{}``` type and methods combined from Go 1 map and slice primitives.

This data structure allows for roughly constant-time reads and for efficiently iterating over only a section of stored values.

```sh
go get -u github.com/umpc/go-sortedmap
```

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
  n := 23
  records := randRecords(n)

  // Create a new collection:
  sm := sortedmap.New(n, asc.Time)

  // The value of n sets the initial backing slice capacity to reduce allocations.
  // Ascending-order sort for time.Time values has been selected for this example.
  // More sort conditional functions are available in this package's subdirectories.

  // Insert the example records:
  sm.BatchInsert(records)

  // Loop through records, in order, until reaching the given value:
  if ch, ok := sm.IterRangeCh(time.Time{}, time.Now()); ok {
    for rec := range ch {
      fmt.Printf("%+v\n", rec)
    }
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
BenchmarkDelete1of1Records-8                 	 5000000	       269 ns/op	       0 B/op	       0 allocs/op

BenchmarkDelete1of10Records-8                	 2000000	       626 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of100Records-8               	 1000000	      1047 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of1000Records-8              	 1000000	      2154 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of10000Records-8             	  300000	      7656 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchDelete10of10Records-8          	  500000	      3307 ns/op	      16 B/op	       1 allocs/op
BenchmarkBatchDelete100of100Records-8        	   30000	     46838 ns/op	     112 B/op	       1 allocs/op
BenchmarkBatchDelete1000of1000Records-8      	    2000	    717205 ns/op	    1024 B/op	       1 allocs/op
BenchmarkBatchDelete10000of10000Records-8    	     100	  19596251 ns/op	   10240 B/op	       1 allocs/op

BenchmarkGet1of1Records-8                    	20000000	        96.6 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchGet10of10Records-8             	 2000000	       904 ns/op	     176 B/op	       2 allocs/op
BenchmarkBatchGet100of100Records-8           	  300000	      5509 ns/op	    1904 B/op	       2 allocs/op
BenchmarkBatchGet1000of1000Records-8         	   30000	     48643 ns/op	   17408 B/op	       2 allocs/op
BenchmarkBatchGet10000of10000Records-8       	    2000	    708432 ns/op	  174080 B/op	       2 allocs/op

BenchmarkHas1of1Records-8                    	20000000	       110 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchHas10of10Records-8             	 2000000	       730 ns/op	      16 B/op	       1 allocs/op
BenchmarkBatchHas100of100Records-8           	  300000	      4934 ns/op	     112 B/op	       1 allocs/op
BenchmarkBatchHas1000of1000Records-8         	   30000	     45088 ns/op	    1024 B/op	       1 allocs/op
BenchmarkBatchHas10000of10000Records-8       	    3000	    513517 ns/op	   10240 B/op	       1 allocs/op

BenchmarkInsert1Record-8                     	 3000000	       431 ns/op	     304 B/op	       2 allocs/op

BenchmarkBatchInsert10Records-8              	  300000	      4160 ns/op	    1382 B/op	       8 allocs/op
BenchmarkBatchInsert100Records-8             	   30000	     54023 ns/op	   14913 B/op	      19 allocs/op
BenchmarkBatchInsert1000Records-8            	    2000	    850022 ns/op	  201972 B/op	      78 allocs/op
BenchmarkBatchInsert10000Records-8           	     100	  22487091 ns/op	 2121013 B/op	     582 allocs/op

BenchmarkReplace1of1Records-8                	 5000000	       368 ns/op	       0 B/op	       0 allocs/op

BenchmarkReplace1of10Records-8               	 1000000	      1006 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplace1of100Records-8              	 1000000	      1676 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplace1of1000Records-8             	  500000	      3261 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplace1of10000Records-8            	  200000	     11758 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchReplace10of10Records-8         	  200000	      9360 ns/op	       0 B/op	       0 allocs/op
BenchmarkBatchReplace100of100Records-8       	   10000	    117045 ns/op	       0 B/op	       0 allocs/op
BenchmarkBatchReplace1000of1000Records-8     	    1000	   1676619 ns/op	       0 B/op	       0 allocs/op
BenchmarkBatchReplace10000of10000Records-8   	      20	  60436417 ns/op	       0 B/op	       0 allocs/op

BenchmarkNew-8                               	20000000	       177 ns/op	      96 B/op	       2 allocs/op
```

The above benchmark tests pause to reset their memory contents between iterations to reduce the interference of CPU caching. The tests were ran on a 4.0GHz Intel Core i7-4790K (Haswell) CPU.

## License

The source code is available under the [MIT License](https://opensource.org/licenses/MIT).
