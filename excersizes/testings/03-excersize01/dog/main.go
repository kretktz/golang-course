// Package dog allows us to more fully understand dogs
package dog

// Years converts human age to dog age
func Years(n int) int {
	return n * 7
}

// YearsTwo converts human age to dog age
func YearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
