package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(5)
	if x != 35 {
		t.Error("Expected 35, got", x)
	}
}

func TestYearsTwo(t *testing.T) {
	x := YearsTwo(5)
	if x != 35 {
		t.Error("Expected 35, got", x)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(5))
	// Output:
	// 35
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(5)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(5)
	}
}
