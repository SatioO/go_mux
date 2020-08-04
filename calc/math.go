package calc

// Add returns sum of two integers
func Add(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}
