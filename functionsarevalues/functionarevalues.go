package functionarevalues

import "fmt"

// Function Types
type transformFn func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	trasnformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTrasnformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(trasnformedNumbers)
	fmt.Println(moreTrasnformedNumbers)
}

// Function as Values
func transformNumbers(numbers *[]int, transorm transformFn) []int {
	tNumbers := []int{}
	for _, value := range *numbers {
		tNumbers = append(tNumbers, transorm(value))
	}
	return tNumbers
}

// Return Function as Values
func getTransformerFunction(number *[]int) transformFn {
	if (*number)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
