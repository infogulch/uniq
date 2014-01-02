package uniq_test

import (
	"github.com/infogulch/uniq"
	"math/rand"
	"sort"
	"testing"
)

type uniquer func(uniq.Interface) int

// Generic test instances below
func TestUniqEmpty(t *testing.T)   { testUniqEmpty(uniq.Uniq, t) }
func TestStableEmpty(t *testing.T) { testUniqEmpty(uniq.Stable, t) }

func TestUniqOne(t *testing.T)   { testUniqOne(uniq.Uniq, t) }
func TestStableOne(t *testing.T) { testUniqOne(uniq.Stable, t) }

func TestUniqTiny(t *testing.T)   { testUniqTiny(uniq.Uniq, t) }
func TestStableTiny(t *testing.T) { testUniqTiny(uniq.Stable, t) }

func TestUniqSmall(t *testing.T)   { testUniqRandom(uniq.Uniq, t, 30, 1000) }
func TestStableSmall(t *testing.T) { testUniqRandom(uniq.Stable, t, 30, 1000) }

func TestUniqMedium(t *testing.T)   { testUniqRandom(uniq.Uniq, t, 1000, 10) }
func TestStableMedium(t *testing.T) { testUniqRandom(uniq.Stable, t, 1000, 10) }

func TestUniqLarge(t *testing.T)   { testUniqRandom(uniq.Uniq, t, 10000, 2) }
func TestStableLarge(t *testing.T) { testUniqRandom(uniq.Stable, t, 10000, 2) }

func TestUniqAllDuplicate(t *testing.T)   { testUniqAllDuplicate(uniq.Uniq, t) }
func TestStableAllDuplicate(t *testing.T) { testUniqAllDuplicate(uniq.Stable, t) }

// Generic tests that apply to all uniquer funcs
func testUniqEmpty(unique uniquer, t *testing.T) {
	a := []int{}
	a = a[:unique(sort.IntSlice(a))]
	if len(a) != 0 {
		t.Errorf("Congratulations, you broke loops")
	}
}

func testUniqOne(unique uniquer, t *testing.T) {
	const value int = 100
	a := []int{value}
	a = a[:unique(sort.IntSlice(a))]
	if len(a) != 1 {
		t.Errorf("Wrong length")
	} else if a[0] != value {
		t.Errorf("Changed values. Should be %d: %d", value, a[0])
	}
}

func testUniqTiny(unique uniquer, t *testing.T) {
	a := []int{1, 2, 2, 2, 3, 4, 4, 4, 4, 4, 5, 5, 6, 6, 7, 8, 9}
	u := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a = a[:unique(sort.IntSlice(a))]
	compareIntSlice(a, u, t)
}

func testUniqRandom(unique uniquer, t *testing.T, size, count int) {
	rand.Seed(0) // specify seed for determinism
	u, a := make([]int, size), make([]int, 100*size)
	for i, _ := range u {
		u[i] = i
	}
	for ; count > 0; count-- {
		for i, _ := range a {
			a[i] = rand.Intn(size)
		}
		sort.Sort(sort.IntSlice(a))
		b := a[:unique(sort.IntSlice(a))]
		compareIntSlice(b, u, t)
	}
}

func testUniqAllDuplicate(unique uniquer, t *testing.T) {
	a := make([]int, 1000)
	const value int = 123
	for i, _ := range a {
		a[i] = value
	}
	a = a[:unique(sort.IntSlice(a))]
	if len(a) != 1 {
		t.Errorf("Didn't eliminate all duplicates. len should be 1: %d", len(a))
	} else if a[0] != value {
		t.Errorf("Changed values. Should be %d: %d", value, a[0])
	}
}

// types for stability tests

// stability tests

// helper functions
func compareIntSlice(a, b []int, t *testing.T) {
	for i, v := range a {
		if v != b[i] {
			t.Errorf("(a[%d] = %d) != (b[%d] = %d)", i, v, i, b[i])
		}
	}
}
