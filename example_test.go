package uniq_test

import (
	"fmt"
	"github.com/infogulch/uniq"
	"sort"
)

func Example() {
	a := []int{1, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9, 9, 9} // already sorted
	a = a[:uniq.Uniq(sort.IntSlice(a))]
	fmt.Println(a)
	// Output: [1 2 3 4 5 6 7 8 9]
}

func Example_unsorted() {
	a := []int{2, 5, 9, 1, 3, 2, 3, 4, 7, 4, 8, 3, 8, 4, 3, 6, 6} // not sorted
	sb := sort.IntSlice(a)
	sort.Sort(sb)         // sort first. doesn't change len
	a = a[:uniq.Uniq(sb)] // reuse the sort.Interface
	fmt.Println(a)
	// Output: [1 2 3 4 5 6 7 8 9]
}
