package util_map

import (
	"fmt"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
)

func TestMapSet(t *testing.T) {
	intSet1 := mapset.NewSet[int](1, 2, 3, 4)
	AddEle(intSet1)
	intSet2 := mapset.NewSet[int](2, 3, 4, 5)

	fmt.Printf("intSet1: %v\n", intSet1) // 1 2 3 4 100
	fmt.Printf("intSet2: %v\n", intSet2) // 2 3 4 5
	// 差集
	diffSet1 := intSet1.Difference(intSet2) // 差集，属于intSet1，不属于intSet2
	diffSet2 := intSet2.Difference(intSet1)
	fmt.Printf("diffSet: %v\n", diffSet1)  // 1 100
	fmt.Printf("diffSet2: %v\n", diffSet2) // 5
	// 交集
	interSet1 := intSet1.Intersect(intSet2)
	interSet2 := intSet2.Intersect(intSet1)
	fmt.Printf("interSet1: %v\n", interSet1) // 2 3 4
	fmt.Printf("interSet2: %v\n", interSet2) // 2 3 4
	// 并集
	unionSet1 := intSet1.Union(intSet2)
	unionSet2 := intSet2.Union(intSet1)
	fmt.Printf("unionSet1: %v\n", unionSet1)                   // 1 2 3 4 5 100
	fmt.Printf("unionSet2.String(): %v\n", unionSet2.String()) // 1 2 3 4 5 100

	// 迭代
	for i := range unionSet1.Iter() {
		fmt.Printf("Iter: i: %v\n", i)
	}

	intSet2.Each(func(i int) bool {
		if i == 3 {
			return true
		}
		i = i * i
		fmt.Printf("Each: i: %v\n", i)
		return false
	})
}

func AddEle(set mapset.Set[int]) {
	set.Add(100)
}
