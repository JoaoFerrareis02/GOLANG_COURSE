package main

import "fmt"

func main() {
	age := 32 // regular variable

	// Creating a pointer
	agePointer := &age

	// Pointers as value
	fmt.Println("Age:", *agePointer)

	// adultYears := getAdultYears(agePointer)
	// fmt.Println(adultYears)

	getAdultYears(agePointer)
	fmt.Println(age)

}

// Using pointers & Passing pointers to functions
// func getAdultYears(age *int) int {
// 	return *age - 18
// }

// Using pointers for data mutation
func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
