package main

import (
	"testing"
)

func Test_Sum(t *testing.T) {
	r := Sum(1, 2, 3)
	if r != 6 {
		t.Error("Expected ", 6, " Got ", r)
	}

	type test struct {
		input  []int
		answer int
	}
	tests := []test{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{5, 6, 7, 8}, 26},
		{[]int{9, 10, 11}, 30},
	}

	for _, v := range tests {
		res := Sum(v.input...)
		if v.answer != res {
			t.Error("Expected: ", v.answer, "Got: ", res)
		}
	}
}
