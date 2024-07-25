package main

import "testing"

func TestLcp(t *testing.T) {
	// 测试用例类型
	type test struct {
		input []string
		want  string
	}

	tests := []test{
		{input: []string{"", ""}, want: ""},
		{input: []string{"", "fd"}, want: ""},
		{input: []string{"fd", ""}, want: ""},
		{input: []string{"jifd", "efeg"}, want: ""},
		{input: []string{"egj", "egwg"}, want: "eg"},
		{input: []string{"yyu", "yyu"}, want: "yyu"},
	}

	for _, tc := range tests {
		got := lcp(tc.input[0], tc.input[1])
		if got != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestLongtestCommonPrefix(t *testing.T) {
	type test struct {
		input []string
		want  string
	}

	tests := []test{
		{input: []string{""}, want: ""},
		{input: []string{"fd"}, want: "fd"},
		{input: []string{"fd", ""}, want: ""},
		{input: []string{"jifd", "efeg"}, want: ""},
		{input: []string{"egj", "egwg"}, want: "eg"},
		{input: []string{"yyu", "yyu"}, want: "yyu"},
		{input: []string{"yyu", "yyu", ""}, want: ""},
		{input: []string{"yyu", "yyu", "yyffdvd"}, want: "yy"},
		{input: []string{"yyu", "fe", "fe"}, want: ""},
		{input: []string{"fegg", "feg", "fedg"}, want: "fe"},
		{input: []string{"eng", "engli", "engi", "engg"}, want: "eng"},
	}

	for _, tc := range tests {
		got := longtestCommonPrefix(tc.input)
		if got != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
