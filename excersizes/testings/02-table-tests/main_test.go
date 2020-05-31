package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		result int
	}

	tests := []test{
		test{[]int{2, 3}, 5},
		test{[]int{2, 3, 4}, 9},
		test{[]int{34, 4}, 38},
		test{[]int{4, 5, 6}, 15},
		test{[]int{2, 9}, 11},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.result {
			t.Error("Expected", v.result, "got", x)
		}
	}

}
