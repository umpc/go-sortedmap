# SortedMap

[![Build Status](https://travis-ci.org/umpc/go-sortedmap.svg?branch=master)](https://travis-ci.org/umpc/go-sortedmap) [![Coverage Status](https://codecov.io/github/umpc/go-sortedmap/badge.svg?branch=master)](https://codecov.io/github/umpc/go-sortedmap?branch=master) [![GoDoc](https://godoc.org/github.com/umpc/go-sortedmap?status.svg)](https://godoc.org/github.com/umpc/go-sortedmap)

SortedMap is a simple library that provides a value-sorted ```map[interface{}]interface{}``` type and methods combined from Go 1 map and slice primitives.

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

  "github.com/umpc/go-sortedmap"
  "github.com/umpc/go-sortedmap/asc"
)

func main() {
  // Create an empty SortedMap with a size suggestion and a less than function:
  sm := sortedmap.New(4, asc.Time)

  // Insert example records:
  sm.Insert("OpenBSD",  time.Date(1995, 10, 18,  8, 37, 1, 0, time.UTC))
  sm.Insert("UnixTime", time.Date(1970,  1,  1,  0,  0, 0, 0, time.UTC))
  sm.Insert("Linux",    time.Date(1991,  8, 25, 20, 57, 8, 0, time.UTC))
  sm.Insert("GitHub",   time.Date(2008,  4, 10,  0,  0, 0, 0, time.UTC))

  // Set iteration options:
  reversed   := true
  lowerBound := time.Date(1994, 1, 1, 0, 0, 0, 0, time.UTC)
  upperBound := time.Now()

  // Select values >= lowerBound and values <= upperBound.
  // Loop through the values, in reverse order:
  if ch, ok := sm.BoundedIterCh(reversed, lowerBound, upperBound); ok {
    for rec := range ch {
      fmt.Printf("%+v\n", rec)
    }
  } else {
    fmt.Println("No values found that were equal to or within the given bounds.")
  }
}
```

Check out the [examples](https://github.com/umpc/go-sortedmap/tree/master/examples), [documentation](https://godoc.org/github.com/umpc/go-sortedmap), and test files, for more features and further explanations.

## Benchmarks

```sh
BenchmarkDelete1of1Records-8                 	 5000000	       285 ns/op	       0 B/op	       0 allocs/op

BenchmarkDelete1of10Records-8                	 2000000	       615 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of100Records-8               	 1000000	      1005 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of1000Records-8              	 1000000	      1987 ns/op	       0 B/op	       0 allocs/op
BenchmarkDelete1of10000Records-8             	  300000	      5473 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchDelete10of10Records-8          	  500000	      3410 ns/op	      16 B/op	       1 allocs/op
BenchmarkBatchDelete100of100Records-8        	   30000	     47069 ns/op	     112 B/op	       1 allocs/op
BenchmarkBatchDelete1000of1000Records-8      	    2000	    721201 ns/op	    1024 B/op	       1 allocs/op
BenchmarkBatchDelete10000of10000Records-8    	     100	  19275331 ns/op	   10240 B/op	       1 allocs/op

BenchmarkGet1of1CachedRecords-8              	50000000	        27.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet1of1Records-8                    	20000000	        96.8 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchGet10of10Records-8             	 2000000	       902 ns/op	     176 B/op	       2 allocs/op
BenchmarkBatchGet100of100Records-8           	  300000	      5550 ns/op	    1904 B/op	       2 allocs/op
BenchmarkBatchGet1000of1000Records-8         	   30000	     49057 ns/op	   17408 B/op	       2 allocs/op
BenchmarkBatchGet10000of10000Records-8       	    2000	    710611 ns/op	  174080 B/op	       2 allocs/op

BenchmarkHas1of1CachedRecords-8              	50000000	        27.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas1of1Records-8                    	20000000	        94.1 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchHas10of10Records-8             	 2000000	       742 ns/op	      16 B/op	       1 allocs/op
BenchmarkBatchHas100of100Records-8           	  300000	      5102 ns/op	     112 B/op	       1 allocs/op
BenchmarkBatchHas1000of1000Records-8         	   30000	     46257 ns/op	    1024 B/op	       1 allocs/op
BenchmarkBatchHas10000of10000Records-8       	    3000	    519497 ns/op	   10240 B/op	       1 allocs/op

BenchmarkInsert1Record-8                     	 3000000	       442 ns/op	     304 B/op	       2 allocs/op

BenchmarkBatchInsert10Records-8              	  300000	      4164 ns/op	    1382 B/op	       8 allocs/op
BenchmarkBatchInsert100Records-8             	   30000	     54184 ns/op	   14912 B/op	      19 allocs/op
BenchmarkBatchInsert1000Records-8            	    2000	    844344 ns/op	  201969 B/op	      78 allocs/op
BenchmarkBatchInsert10000Records-8           	     100	  25911455 ns/op	 2121554 B/op	     584 allocs/op

BenchmarkReplace1of1Records-8                	 5000000	       378 ns/op	       0 B/op	       0 allocs/op

BenchmarkReplace1of10Records-8               	 1000000	      1021 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplace1of100Records-8              	 1000000	      1669 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplace1of1000Records-8             	  500000	      3187 ns/op	       0 B/op	       0 allocs/op
BenchmarkReplace1of10000Records-8            	  200000	      8778 ns/op	       0 B/op	       0 allocs/op

BenchmarkBatchReplace10of10Records-8         	  200000	      6934 ns/op	       0 B/op	       0 allocs/op
BenchmarkBatchReplace100of100Records-8       	   10000	    100186 ns/op	       0 B/op	       0 allocs/op
BenchmarkBatchReplace1000of1000Records-8     	    1000	   1546355 ns/op	       0 B/op	       0 allocs/op
BenchmarkBatchReplace10000of10000Records-8   	      20	  58396360 ns/op	       0 B/op	       0 allocs/op

BenchmarkNew-8                               	20000000	        98.8 ns/op	      96 B/op	       2 allocs/op
```

The above benchmark tests were ran using a 4.0GHz Intel Core i7-4790K (Haswell) CPU.

## License

The source code is available under the [MIT License](https://opensource.org/licenses/MIT).
