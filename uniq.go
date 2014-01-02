// Package uniq provides primitives for getting the first unique elements of slices or user-defined collections from an already sorted list using your existing sort.Interface
package uniq

import "sort"

type Interface sort.Interface

// Uniq moves the unique elements from a sorted sort.Interface to the beginning of the collection, while keeping them in sorted order, and returns the length of the portion of the collection that has unique elements.
// The area where the duplicates are stored is no longer sorted.
func Uniq(data Interface) int {
	len := data.Len()
	if len <= 1 {
		return len
	}
	i, j := 0, 1
	// find the first duplicate
	for j < len && data.Less(i, j) {
		i++
		j++
	}
	// this loop is simpler after the first duplicate is found
	for ; j < len; j++ {
		if data.Less(i, j) {
			i++
			data.Swap(i, j)
		}
	}
	return i + 1
}
