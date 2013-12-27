package uniq

import "sort"

type Interface sort.Interface

func Ue(data Interface) int {
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
