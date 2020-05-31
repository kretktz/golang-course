package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{234, 56443, 245, 25, 36, 45745, 346, 25, 2453})
	}
}

func ExampleCenteredAvg() {
	xi := []int{1, 2, 3, 4, 5}
	a := CenteredAvg(xi)
	fmt.Println(a)
	// Output
	// 3
}

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		result float64
	}

	tests := []test{
		test{[]int{1, 2, 3, 4, 5}, 3},
		test{[]int{2, 3, 4}, 3},
		test{[]int{15, 16, 17, 18, 19}, 17},
		test{[]int{4, 5, 6}, 5},
		test{[]int{2, 9, 18, 27, 567}, 18},
	}

	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.result {
			t.Error("Expected", v.result, "got", x)
		}
	}
}
