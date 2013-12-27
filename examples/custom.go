package main

import (
	"fmt"
	"github.com/infogulch/uniq"
	"math/rand"
	"sort"
)

type Point struct {
	X, Y int
}

// Only one Tag allowed per Point
type PointData struct {
	Point
	Tag string
}

// type implements sort.Interface. See sort for details of this pattern
type PointsData []PointData

func (p PointsData) Len() int      { return len(p) }
func (p PointsData) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PointsData) Less(i, j int) bool {
	// sort by X and then Y, ignore Tag
	return p[i].X < p[j].X || (p[i].X == p[j].X && p[i].Y < p[j].Y)
}

// randomly generate `count` points with X, Y in [0,rang) with Tag tag
func GenPoints(count, rang int, tag string) []PointData {
	ret := make([]PointData, count)
	for i, _ := range ret {
		ret[i] = PointData{Point{rand.Intn(rang), rand.Intn(rang)}, tag}
	}
	return ret
}

func main() {
	// we want all unique a's, and fill in extra space with b's
	a := GenPoints(7, 3, "a")
	b := GenPoints(7, 3, "b")
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	a = append(a, b...)
	as := PointsData(a)
	sort.Stable(as)     // stable sort so the a's are first
	a = a[:uniq.Ue(as)] // discard dup a's and b's
	fmt.Printf("after: %v\n", a)
}
