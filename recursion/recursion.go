package recursion

import "fmt"

func main() {

	fact := factorial(5)
	fmt.Println(fact)

}

// Making Sense of Recursion
func factorial(number int) int {

	if number > 1 {
		return number * factorial(number-1)
	} else {
		return 1
	}

	// result := 1
	// for i := 1; i <= number; i++ {
	// 	result *= i
	// }
	// return result
}
