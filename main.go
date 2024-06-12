package main

import "fmt"

// Working with type alias
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	// "Make" Function - make([]type, length, capacity)
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Max")
	userNames = append(userNames, "João")

	fmt.Println(userNames, len(userNames), cap(userNames))

	// "Make"ing Maps - make(map[type]type, capacity)
	courses := make(floatMap, 3)

	courses["Go"] = 4.7
	courses["React"] = 4.8
	courses["Angular"] = 4.7

	courses.output()
	// fmt.Println(courses)

	// For Loops with Arrays, Slices & Maps

	// Loop of Arrays and Slices
	for index, value := range userNames {
		fmt.Println(index, "-", value)
	}

	// Loop of Maps
	for key, value := range courses {
		fmt.Println(key, "-", value)
	}

}
