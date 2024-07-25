package main

import "fmt"

type User struct {
	Name   string
	Age    uint8
	Gender int8 // 0|女 1|男
}

// func NewUser(name string, age uint8, gender int8) *User {
// 	return &User{
// 		Name:   name,
// 		Age:    age,
// 		Gender: gender,
// 	}
// }

// 定义接收*User作为参数并且会在函数内部修改其字段的函数
type OptionFunc func(*User)

func WithName(name string) OptionFunc {
	return func(u *User) {
		u.Name = name
	}
}

func WithAge(age uint8) OptionFunc {
	return func(u *User) {
		u.Age = age
	}
}

func WithGender(gender int8) OptionFunc {
	return func(u *User) {
		u.Gender = gender
	}
}

func NewUser(name string, opts ...OptionFunc) *User {
	user := &User{Name: name}
	for _, opt := range opts {
		opt(user)
	}
	return user
}

func main() {
	user := NewUser("Jerry", WithAge(20), WithGender(1))
	fmt.Println(user.Age, user.Gender, user.Name)
}
