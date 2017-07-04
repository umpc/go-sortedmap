# SortedMap Examples

## Table of Contents

* [Test Data](#test-data)
* [Insert / Get / Replace / Delete / Has](#insert--get--replace--delete--has)
* [Iteration](#iteration)
  *  [IterCh](#iterch)
  *  [BoundedIterCh](#boundediterch)
  *  [CustomIterCh](#customiterch)
  *  [IterFunc](#iterfunc)
  *  [BoundedIterFunc](#boundediterfunc)
  *  [Map & Keys Loop](#map--keys-loop)
  *  [Map & Bounded Keys Loop](#map--bounded-keys-loop)
  *  [Bounded Delete](#boundeddelete)

## Test Data

The following function is used to generate test data in the examples:

```go
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

## Insert / Get / Replace / Delete / Has

Below is an example containing common operations that are used with a single record.

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
  const n = 1

  records := randRecords(n)
  rec := records[0]

  // Create a new collection. This reserves memory for one item
  // before allocating a new backing array and appending to it:
  sm := sortedmap.New(n, asc.Time)

  // Insert the example record:
  if !sm.Insert(rec.Key, rec.Val) {
    fmt.Printf("The key already exists: %+v", rec.Key)
  }

  // Get and print the value:
  if val, ok := sm.Get(rec.Key); ok {
    fmt.Printf("%+v\n", val)
  }

  // Replace the example record:
  sm.Replace(rec.Key, time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC))

  // Remove the example record:
  sm.Delete(rec.Key)

  // Check if the index has the key:
  const recHasFmt = "The key %v.\n"
  if sm.Has(rec.Key) {
    fmt.Printf(recHasFmt, "exists")
  } else {
    fmt.Printf(recHasFmt, "does not exist")
  }
}
```

## Iteration

SortedMap supports three specific ways of processing iterable data: 

* Channels
* Callback Functions
* Maps & Slices

### IterCh

```IterCh``` is a simple way of iterating over the entire set, in order.

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
  const n = 25
  records := randRecords(n)

  // Create a new collection.
  sm := sortedmap.New(n, asc.Time)

  // BatchInsert the example records:
  sm.BatchInsert(records)

  iterCh := sm.IterCh()
  defer iterCh.Close()

  for rec := range iterCh.Records() {
    fmt.Printf("%+v\n", rec)
  }
}
```

### BoundedIterCh

```BoundedIterCh``` selects the records equal to or between the given bounds. Its first argument allows for reversing the order of the returned records.

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
  const n = 25
  records := randRecords(n)

  // Create a new collection.
  sm := sortedmap.New(n, asc.Time)

  // BatchInsert the example records:
  sm.BatchInsert(records)

  iterCh, err := sm.BoundedIterCh(false, time.Time{}, time.Now())
  if err != nil {
    fmt.Println(err)
  }
  defer iterCh.Close()

  for rec := range iterCh.Records() {
    fmt.Printf("%+v\n", rec)
  }
}
```

### CustomIterCh

```CustomIterCh``` provides all ```IterCh``` functionality and accepts a special ```IterChParams``` value to read settings.

This method offers a deadline timeout for channel sends and it is *highly recommended* that this method, or Map + Slice, be used where reliability over a long run-time counts.

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
  const n = 25
  records := randRecords(n)

  // Create a new collection.
  sm := sortedmap.New(n, asc.Time)

  // BatchInsert the example records:
  sm.BatchInsert(records)

  params := sortedmap.IterChParams{
    SendTimeout: 5 * time.Minute,
    Reversed: true,
  }

  iterCh, err := sm.CustomIterCh(false, time.Time{}, time.Now())
  if err != nil {
    fmt.Println(err)
  }
  defer iterCh.Close()

  for rec := range iterCh.Records() {
    fmt.Printf("%+v\n", rec)
  }
}
```

### IterFunc

```IterFunc``` is like ```IterCh```, but runs a callback function passing in each record instead of sending the records on a channel.

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
  const n = 25
  records := randRecords(n)

  // Create a new collection.
  sm := sortedmap.New(n, asc.Time)

  // BatchInsert the example records:
  sm.BatchInsert(records)

  sm.IterFunc(false, func(rec sortedmap.Record) bool {
    fmt.Printf("%+v\n", rec)
    return true
  })
}
```

### BoundedIterFunc

```BoundedIterFunc``` is the callback function equivalent to ```BoundedIterCh```/```CustomIterCh```.

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
  const n = 25
  records := randRecords(n)

  // Create a new collection.
  sm := sortedmap.New(n, asc.Time)

  // BatchInsert the example records:
  sm.BatchInsert(records)

  if !sm.BoundedIterFunc(false, time.Time{}, time.Now(), func(rec sortedmap.Record) bool {
    fmt.Printf("%+v\n", rec)
    return true
  }) {
    fmt.Println("No values fall within the specified bounds.")
  }
}
```

### Map & Keys Loop

The ```Map``` and ```Keys``` methods offer a way of iterating throughout the map using a combination of Go's native map and slice types.

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
  const n = 25
  records := randRecords(n)

  // Create a new collection.
  sm := sortedmap.New(n, asc.Time)

  // BatchInsert the example records:
  sm.BatchInsert(records)

  m, keys := sm.Map(), sm.Keys()
  for _, k := range keys {
    fmt.Printf("%+v\n", m[k])
  }
}
```

### Map & Bounded Keys Loop

Like the above ```Bounded``` methods, this method allows for selecting only the necessary data for processing.

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
  const n = 25
  records := randRecords(n)

  // Create a new collection.
  sm := sortedmap.New(n, asc.Time)

  // BatchInsert the example records:
  sm.BatchInsert(records)

  // Copy the map + slice headers.
  m := sm.Map()
  if keys, ok := sm.BoundedKeys(time.Time{}, time.Now()); ok {
    for _, k := range keys {
      fmt.Printf("%+v\n", m[k])
    }
  } else {
    fmt.Println("No values fall within the specified bounds.")
  }
}
```

### BoundedDelete

```BoundedDelete``` is a similar pattern as the above ```Bounded``` methods. ```BoundedDelete``` removes values that are equal to or between the provided bounds values.

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
  const n = 25
  records := randRecords(n)

  // Create a new collection.
  sm := sortedmap.New(n, asc.Time)

  // BatchInsert the example records:
  sm.BatchInsert(records)

  // Delete values equal to or within the given bound values.
  if err := sm.BoundedDelete(time.Time{}, time.Now()); err != nil {
    fmt.Println(err)
  }
}
```

For more options and features, check out the documentation:

[![GoDoc](https://godoc.org/github.com/umpc/go-sortedmap?status.svg)](https://godoc.org/github.com/umpc/go-sortedmap)