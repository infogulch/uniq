# uniq
    import "github.com/infogulch/uniq"

Package uniq provides primitives for getting the first unique elements of slices
or user-defined collections from an *already sorted* list using your existing
sort.Interface.

#### Example
    a := []int{1, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9, 9, 9} // already sorted
    a = a[:uniq.Ints(a)]
    fmt.Println(a)

Output: `[1 2 3 4 5 6 7 8 9]`

The unique functions return the length of the unique portion of the collection,
and using the return value as the start or end index in a slice operation (as in
the example above) makes it easy to get the result you need when using slices.

*Why use this instead of a map?* uniq doesn't need extra memory. If you're
sorting already or the result needs to be sorted uniq is faster. Uniq is much
easier to use and implement in your code, and more easily extended.

*Why use a sort.Interface?* Because the algorithms require the data be sorted
beforehand anyway, so you don't have to have to implement anything else, and
because they don't need any more. `Uniq` is O(n) and probably couldn't be
much faster even considering the sort step without using more memory. `Stable`
is slower and I've considered extending the interface to support copying
elements to a duplicates buffer but nobody has expressed a need.

--
#### Index
* [func Uniq(data Interface) int](#func-uniq)
* [func Stable(data Interface) int](#func-stable)
* [func IsUnique(data Interface) bool](#func-isunique)
* [func Float64s(a []float64) int](#func-float64s)
* [func Float64sAreUnique(a []float64) bool](#func-float64sareunique)
* [func Ints(a []int) int](#func-ints)
* [func IntsAreUnique(a []int) bool](#func-intsareunique)
* [func Strings(a []string) int](#func-strings)
* [func StringsAreUnique(a []string) bool](#func-stringsareunique)
* [type Interface](#type-interface)

--
#### func Uniq
    func Uniq(data Interface) int

Uniq moves the first unique elements to the beginning of the *sorted* collection
and returns the number of unique elements.

It makes one call to data.Len to determine n, n-1 calls to data.Less, and O(n)
calls to data.Swap. The unique elements remain in original sorted order, but the
duplicate elements do not.

--
#### func Stable
    func Stable(data Interface) int

Stable moves the first unique elements to the beginning of the *sorted*
collection and returns the number of unique elements, but also keeps the
original order of duplicate elements.

It makes one call to data.Len, O(n) calls to data.Less, and O(n*log(n)) calls to
data.Swap.

--
#### func IsUnique
    func IsUnique(data Interface) bool

IsUnique reports whether data is sorted and unique.

--
#### func Float64s
    func Float64s(a []float64) int

Float64s calls unique on a slice of float64.

--
#### func Float64sAreUnique
    func Float64sAreUnique(a []float64) bool

Float64sAreUnique tests whether the slice of float64 is sorted and unique.

--
#### func Ints
    func Ints(a []int) int

Ints calls unique on a slice of int.

--
#### func IntsAreUnique
    func IntsAreUnique(a []int) bool

IntsAreUnique tests whether the slice of int is sorted and unique.

--
#### func Strings
    func Strings(a []string) int

Strings calls unique on a slice of string.

--
#### func StringsAreUnique
    func StringsAreUnique(a []string) bool

StringsAreUnique tests whether the slice of string is sorted and unique.

--
#### type Interface
    type Interface interface {
        // Len returns the number of elements.
        Len() int
        // Less tells if the element at index i should come
        // before the element at index j.
        Less(i, j int) bool
        // Swap swaps the elements at indexes i and j.
        Swap(i, j int)
    }

Interface to use the uniq package. Identical to sort.Interface.
