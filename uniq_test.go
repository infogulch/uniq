package uniq

import "math/rand"
import "sort"
import "testing"

func TestUniqueEmpty(t *testing.T) {
	a := []int{}
	a = a[:Ue(sort.IntSlice(a))]
	if len(a) != 0 {
		t.Errorf("Congratulations, you broke loops")
	}
}

func TestUniqueOne(t *testing.T) {
	const value int = 100
	a := []int{value}
	a = a[:Ue(sort.IntSlice(a))]
	if len(a) != 1 {
		t.Errorf("Wrong length")
	} else if a[0] != value {
		t.Errorf("Changed values. Should be %d: %d", value, a[0])
	}
}

func TestAllDuplicate(t *testing.T) {
	a := make([]int, 1000)
	const value int = 123
	for i, _ := range a {
		a[i] = value
	}
	a = a[:Ue(sort.IntSlice(a))]
	if len(a) != 1 {
		t.Errorf("Didn't eliminate all duplicates. len should be 1: %d", len(a))
	} else if a[0] != value {
		t.Errorf("Changed values. Should be %d: %d", value, a[0])
	}
}

func TestUniqueSmall(t *testing.T) {
	a := []int{1, 2, 2, 2, 3, 4, 4, 4, 4, 4, 5, 5, 6, 6, 7, 8, 9}
	u := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a = a[:Ue(sort.IntSlice(a))]
	compareIntSlice(a, u, t)
}

func TestUniqueLargeRand(t *testing.T) {
	rand.Seed(0) // specify seed for determinism
	const size int = 1000
	u := make([]int, size)
	for i, _ := range u {
		u[i] = i
	}
	for n := 0; n < 10; n++ {
		a := make([]int, 100*size)
		for i, _ := range a {
			a[i] = rand.Intn(size)
		}
		sort.Sort(sort.IntSlice(a))
		a = a[:Ue(sort.IntSlice(a))]
		compareIntSlice(a, u, t)
	}
}

func compareIntSlice(a, b []int, t *testing.T) {
	for i, v := range a {
		if v != b[i] {
			t.Errorf("(a[%d] = %d) != (b[%d] = %d)", i, v, i, b[i])
		}
	}
}
