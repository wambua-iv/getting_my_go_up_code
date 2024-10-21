package main

import (
	"fmt"
)

// isPowerOfTwo checks if a given number is a power of 2
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 8, 16, 18, 32, 64, 100}

	for _, num := range numbers {
		if isPowerOfTwo(num) {
			fmt.Printf("%d is a power of 2\n", num)
		} else {
			fmt.Printf("%d is not a power of 2\n", num)
		}
	}
}
