uniq
====

A Golang package to get the first unique elements of slices or user-defined
collections from an already sorted list using your existing sort.Interface

Usage
-----

```
uniq
unique
uniq.ue
uniq.Ue
uniq.Ue(sort.Interface) (unique_length int)
```

The list *must* already be sorted, otherwise the behavior is undefined. Since
you just sorted it with `sort.Sort` (you did, right?) you can call `uniq.Ue`
with the `sort.Interface` that you just created, because uniq uses it verbatim.
`uniq.Ue` will call `Len` once, `Less` up to `n+1` times, and `Swap` up to `n-1`
times and return the length of the unique portion of the array. The remaining
part of the array will contain all the duplicates but is no longer sorted. You
are still responsible for any cleanup of the duplicates if necessary.

If you're using a slice, re-slice from the beginning and use the return result
from `uniq.Ue` as the end index: `arr = arr[:uniq.Ue(ArrSorter(arr))]`

If you need to do cleanup of duplicates, I suggest making a `func Cleanup([]T)`,
then you can do this:

```go
n := uniq.Ue(ArrSorter(arr))
Cleanup(arr[n:]) // this is safe even if n == len(arr)
arr = arr[:n]
```

Examples
--------

Here are some simple examples. This and others are go run-able from the 
[examples](examples) dir.

```go
package main

import (
	"fmt"
	"github.com/infogulch/uniq"
	"sort"
)

func main() {
	// already sorted
	a := []int{1, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9, 9, 9}
	a = a[:uniq.Ue(sort.IntSlice(a))]
	fmt.Println(a)

	// unsorted
	b := []int{2, 5, 9, 1, 3, 2, 3, 4, 7, 4, 8, 3, 8, 4, 3, 6, 6}
	sb := sort.IntSlice(b)
	sort.Sort(sb)       // sort first. doesn't change len
	b = b[:uniq.Ue(sb)] // reuse the sort.Interface
	fmt.Println(b)
}
```
