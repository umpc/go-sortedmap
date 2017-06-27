# SortedMap

[![Build Status](https://travis-ci.org/umpc/go-sortedmap.svg?branch=master)](https://travis-ci.org/umpc/go-sortedmap) [![Coverage Status](https://codecov.io/github/umpc/go-sortedmap/badge.svg?branch=master)](https://codecov.io/github/umpc/go-sortedmap?branch=master) [![GoDoc](https://godoc.org/github.com/umpc/go-sortedmap?status.svg)](https://godoc.org/github.com/umpc/go-sortedmap)

SortedMap is a small library that provides a value-sorted ```map[interface{}]interface{}``` type and methods combined from Go 1 map and slice primitives.

This data structure allows for roughly constant-time reads and for efficiently iterating over only a subsection of the stored values. While this works for the project it was built for, because of the structure's reliance on a sorted slice of keys, worst-case delete operations are roughly ```O(n)```, where ```n``` is the number of items in the collection.

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

## License

The source code is available under the [MIT License](https://opensource.org/licenses/MIT).
