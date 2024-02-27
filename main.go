package main

func main() {
	// Generate the next 5 numbers in the Fibonacci sequence
	fibonacciSequence(0)
}

func fibonacciSequence(start int) []int {
	// Find the starting point in the Fibonacci sequence
	a, b := 0, 1
	for b < start {
		a, b = b, a+b
	}

	// Generate the next 5 numbers in the sequence
	result := make([]int, 5)
	for i := 0; i < 5; i++ {
		a, b = b, a+b
		result[i] = a
	}

	return result
}
