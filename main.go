package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	// Spliting Slices Into Parameter Values
	sum := sumup(numbers...)

	fmt.Println(sum)

}

// Variadic Functions - Do not force to pass arrays for values
func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
