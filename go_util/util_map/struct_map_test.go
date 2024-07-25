package util_map_test

import (
	"fmt"
	"go_util/util_map"
	"testing"
)

// 示例结构体
type Address struct {
	City    string
	Country string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

func TestStructToMap(t *testing.T) {
	user := User{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:    "Wonderland",
			Country: "Fantasy",
		},
	}
	result := util_map.StructToMap(user)
	fmt.Printf("result: %+v\n", result)
}
