package uniq_test

import (
	"github.com/infogulch/uniq"
	"math/rand"
	"sort"
	"testing"
)

type uniquer func(uniq.Interface) int

func testUniqEmpty(unique uniquer, t *testing.T) {
	a := []int{}
	a = a[:unique(sort.IntSlice(a))]
	if len(a) != 0 {
		t.Errorf("Congratulations, you broke loops")
	}
}

func TestUniqEmpty(t *testing.T)   { testUniqEmpty(uniq.Uniq, t) }
func TestStableEmpty(t *testing.T) { testUniqEmpty(uniq.Stable, t) }

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

func TestUniqOne(t *testing.T)   { testUniqOne(uniq.Uniq, t) }
func TestStableOne(t *testing.T) { testUniqOne(uniq.Stable, t) }

func testUniqTiny(unique uniquer, t *testing.T) {
	a := []int{1, 2, 2, 2, 3, 4, 4, 4, 4, 4, 5, 5, 6, 6, 7, 8, 9}
	u := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a = a[:unique(sort.IntSlice(a))]
	compareIntSlice(a, u, t)
}

func TestUniqTiny(t *testing.T)   { testUniqTiny(uniq.Uniq, t) }
func TestStableTiny(t *testing.T) { testUniqTiny(uniq.Stable, t) }

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
		sort.Ints(a)
		b := a[:unique(sort.IntSlice(a))]
		compareIntSlice(b, u, t)
	}
}

func TestUniqSmall(t *testing.T)   { testUniqRandom(uniq.Uniq, t, 30, 1000) }
func TestStableSmall(t *testing.T) { testUniqRandom(uniq.Stable, t, 30, 1000) }

func TestUniqMedium(t *testing.T)   { testUniqRandom(uniq.Uniq, t, 1000, 10) }
func TestStableMedium(t *testing.T) { testUniqRandom(uniq.Stable, t, 1000, 10) }

func TestUniqLarge(t *testing.T)   { testUniqRandom(uniq.Uniq, t, 10000, 2) }
func TestStableLarge(t *testing.T) { testUniqRandom(uniq.Stable, t, 10000, 2) }

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

func TestUniqAllDuplicate(t *testing.T)   { testUniqAllDuplicate(uniq.Uniq, t) }
func TestStableAllDuplicate(t *testing.T) { testUniqAllDuplicate(uniq.Stable, t) }

// types for stability tests
type Pair struct{ a, b int }
type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pairs) Less(i, j int) bool { return p[i].a < p[j].a }

func TestStability(t *testing.T) {
	n := 10000
	a := make([]Pair, n)
	for i, _ := range a {
		a[i] = Pair{rand.Intn(100), i}
	}
	sort.Stable(Pairs(a))
	r := uniq.Stable(Pairs(a))
	if !uniq.IsUnique(Pairs(a[:r])) {
		t.Errorf("Not unique")
	}
	if !sort.IsSorted(Pairs(a[r:])) {
		t.Errorf("Not sorted")
	}
	for i := 1; i < n; i++ {
		if a[i-1].a == a[i].a && a[i-1].b >= a[i].b {
			t.Errorf("Not stable")
		}
	}
}

// helper functions
func compareIntSlice(a, b []int, t *testing.T) {
	for i, v := range a {
		if v != b[i] {
			t.Errorf("(a[%d] = %d) != (b[%d] = %d)", i, v, i, b[i])
		}
	}
}
