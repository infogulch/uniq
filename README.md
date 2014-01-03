# uniq

    import "github.com/infogulch/uniq"

Package uniq provides primitives for getting the first unique elements of slices
or user-defined collections from an *already sorted* list using your existing
sort.Interface.

#### Index
* [func Uniq(data Interface) int](#func-uniq)
* [func Stable(data Interface) int](#func-stable)
* [func IsSortedUnique(data Interface) bool](#func-issortedunique)
* [type Interface](#type-interface)

#### Example
```go
a := []int{1, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9, 9, 9} // already sorted
a = a[:uniq.Uniq(sort.IntSlice(a))]
fmt.Println(a)
```

    [1 2 3 4 5 6 7 8 9]

## Usage

#### func Uniq

```go
func Uniq(data Interface) int
```
Uniq moves the first unique elements to the beginning of the *sorted* collection
and returns the number of unique elements.

It makes one call to data.Len to determine n, n-1 calls to data.Less, and O(n)
calls to data.Swap. The unique elements remain in original sorted order, but the
duplicate elements do not.

--

#### func Stable

```go
func Stable(data Interface) int
```
Stable moves the first unique elements to the beginning of the *sorted*
collection and returns the number of unique elements, but also keeps the
original order of duplicate elements.

It makes one call to data.Len, O(n) calls to data.Less, and O(n*log(n)) calls to
data.Swap.

--

#### func IsSortedUnique

```go
func IsSortedUnique(data Interface) bool
```
IsSortedUnique reports whether data is sorted and unique.

--

#### type Interface

```go
type Interface interface {
	// Len returns the number of elements.
	Len() int
	// Less tells if the element at index i should come
	// before the element at index j.
	Less(i, j int) bool
	// Swap swaps the elements at indexes i and j.
	Swap(i, j int)
}
```
Identical to the sort package interface.
