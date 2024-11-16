package main

import (
	"fmt"
)

/**
	Sum and average of a slice of integers
*/
func sum(vals []int) int {
	sum := 0
	for _, val := range vals {
		sum += val
	}
	return sum
}

/**
	Average of a slice of integers
*/
func avg(vals []int) float64 {
	return float64(sum(vals)) / float64(len(vals))
}

func main() {
	var numbers = [...]int{1, 2, 3, 4, 5, 8, 99, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Sum: %d\n", sum(numbers[:]))
	fmt.Printf("Average: %0.3f\n", avg(numbers[:]))
}
