package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {

	// 1)
	hobbies := [3]string{"Hiking", "Walking", "Biking"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3)
	selectedHobbies1 := hobbies[0:2]
	selectedHobbies2 := hobbies[:2]

	fmt.Println(selectedHobbies1)
	fmt.Println(selectedHobbies2)

	// 4)
	selectedHobbies3 := selectedHobbies2[1:3]
	fmt.Println(selectedHobbies3)

	// 5)
	goals := []string{"Be the very best!", "Eat hamburger!"}
	fmt.Println(goals)

	// 6)
	goals[1] = "Be gentle!"
	goals = append(goals, "Be grateful!")
	fmt.Println(goals)

	// 7)
	products := []Product{
		{
			title: "X - Tudo",
			id:    1,
			price: 25.99,
		},
		{
			title: "X - Egg Bacon",
			id:    2,
			price: 15.99,
		},
	}

	newProduct := Product{
		title: "X - Bacon",
		id:    3,
		price: 12.99,
	}

	products = append(products, newProduct)

	fmt.Println(products)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
