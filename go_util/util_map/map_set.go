package util_map

import (
	mapset "github.com/deckarep/golang-set/v2"
)

func MapSet() {
	intSet := mapset.NewSet[int]()
	intSet.Add(1)
	intSet.Append([]int{9, 0, -1}...)
}
