package common

import "golang.org/x/exp/constraints"

type _NilType struct{}

var _Nil = _NilType{}

type Set[T constraints.Ordered] map[T]_NilType

func NewSet[T constraints.Ordered](keys ...T) Set[T] {
	m := make(Set[T], len(keys))
	for _, key := range keys {
		m[key] = _Nil
	}
	return m
}
func (m Set[T]) Add(keys ...T) {
	m.Adds(keys)
}
func (m Set[T]) Adds(keys []T) {
	for _, key := range keys {
		m[key] = _Nil
	}
}

func (m Set[T]) In(key T) bool {
	_, ok := m[key]
	return ok
}

func (m Set[T]) InOrAdd(key T) bool {
	_, ok := m[key]
	if !ok {
		m[key] = _Nil
	}
	return ok
}
