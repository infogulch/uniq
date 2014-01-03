package uniq_test

import (
	"fmt"
	"github.com/infogulch/uniq"
	"sort"
)

type Item struct {
	Name  string
	Price float32
}

// Implements sort.Interface
type Items []Item

func (p Items) Len() int           { return len(p) }
func (p Items) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Items) Less(i, j int) bool { return p[i].Name < p[j].Name }

func ExampleUniq_interface() {
	olditems := []Item{{"Apple", 0.55}, {"Orange", 0.82}, {"Banana", 0.37}}
	newitems := []Item{{"Apple", 0.50}, {"Orange", 1.02}, {"Peach", 1.57}}
	newitems = append(newitems, olditems...)
	sort.Stable(Items(newitems))
	newitems = newitems[:uniq.Uniq(Items(newitems))]
	fmt.Println(newitems)

	// Output: [{Apple 0.5} {Banana 0.37} {Orange 1.02} {Peach 1.57}]
}
