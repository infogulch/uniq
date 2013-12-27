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
