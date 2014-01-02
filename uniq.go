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

func Stable(data Interface) int {
	return stable(data, 0, data.Len())
}

func stable(data Interface, start, end int) int {
	if n := end - start; n <= 2 {
		if n == 2 && !data.Less(start, start+1) {
			n--
		}
		return n
	}
	mid := start + (end-start)/2 // average safe from overflow
	ua := stable(data, start, mid)
	ub := stable(data, mid, end)
	if ua > 0 && ub > 0 && !data.Less(start+ua-1, mid) {
		mid++ // the first element in B is present in a
		ub--
	}
	shift(data, start+ua, mid, mid+ub)
	return ua + ub
}

// shift exchanges elements in a sort.Interface from range [start,mid) with those in range [mid,end).
// Performs (end-start) swaps and accesses each element twice in the worst & average case, and (end-start)/2 swaps and accesses each element once in the best case.
func shift(data Interface, start, mid, end int) {
	if start >= mid || mid >= end {
		return // no elements to shift
	}
	if mid-start == end-mid {
		// equal sizes, use faster algorithm
		swapn(data, start, mid, mid-start)
		return
	}
	reverse(data, start, mid)
	reverse(data, mid, end)
	reverse(data, start, end)
}

// reverse transposes elements in a sort.Interface so that the elements in range [start,end) are in reverse order.
// Performs (end-start)/2 swaps, and acceses each element from start to end once.
func reverse(data Interface, start, end int) {
	end--
	for start < end {
		data.Swap(start, end)
		start++
		end--
	}
}

// swapn transposes the elements in two sections of equal length in a sort.Interface.
// The sections start at indices i & j.
// If the sections overlap (i.e. min(i,j)+len > max(i,j)) the result is undefined.
// Performs len swaps and acceses each element once.
func swapn(data Interface, i, j, n int) {
	for n > 0 {
		n--
		data.Swap(i+n, j+n)
	}
}
