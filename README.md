uniq
====

A Golang package to get the first unique elements of slices or user-defined collections from an already sorted list. Uses your existing sort.Interface

Usage
=====

```
uniq
unique
uniq.ue
uniq.Ue
uniq.Ue(sort.Interface) (unique_length int)
```

The list *must* already be sorted, otherwise the behavior is undefined. Since you just sorted it with `sort.Sort` (you did, right?) you can call `uniq.Ue` with the `sort.Interface` that you just created, because uniq uses it verbatim. `uniq.Ue` will call `Len` once, `Less` up to `n+1` times, and `Swap` up to `n-1` times and return the length of the unique portion of the array. The remaining part of the array will contain all the duplicates but is no longer sorted.

If you're using a slice, re-slice from the beginning and use the return result from `uniq.Ue` as the end index: `arr = arr[:uniq.Ue(ArrSorter(arr))]`

Examples
========

Simple examples:
```go
package main

import (
  "github.com/infogulch/uniq"
  "sort"
  "fmt"
}

func main() {
  // already sorted
  a := []int{1,2,3,3,3,4,5,6,7,8,9,9,9}
  a = a[:uniq.Ue(sort.IntSlice(a))]
  fmt.Println(a)
  
  // unsorted
  b := []int{2,5,1,3,2,3,4,4,8,3,8,4,3,6,6}
  sb := sort.IntSlice(b)
  sort.Sort(sb)       // sort first. doesn't change len
  b = b[:uniq.Ue(sb)] // reuse the sort.Interface
  fmt.Println(b)
}
```
