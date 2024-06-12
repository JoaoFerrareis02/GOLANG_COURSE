package lists

// import "fmt"

// func main() {

// // Introducing Arrays
// prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// fmt.Println(prices)

// //Working with Arrays
// prices[2] = 32.99      // Changing an element
// fmt.Println(prices[2]) // Get an item

// // Selecting Parts of Arrays with Slices
// featuredPrices := prices[1:3] // Selecting the 1º and 2º element of prices
// fmt.Println(featuredPrices)

// featuredPrices = prices[:3] // Selecting from the beggining to the 2º element of prices
// fmt.Println(featuredPrices)

// featuredPrices = prices[1:] // Selecting the 1º to the last element of prices
// fmt.Println(featuredPrices)

// highligthedPrices := featuredPrices[:1] // Creating a Slice of an Slice
// fmt.Println(highligthedPrices)

// // Diving depper into Slices

// featuredPrices[0] = 199.99 // Slices works like pointers to arrays
// fmt.Println(prices)

// fmt.Println(len(featuredPrices), cap(featuredPrices)) // Lenght & Capacity of an Slice
// fmt.Println(len(highligthedPrices), cap(highligthedPrices))

// highligthedPrices = highligthedPrices[:3]
// fmt.Println(highligthedPrices)
// fmt.Println(len(highligthedPrices), cap(highligthedPrices))

// 	// Slices can be used as dynamic arrays
// 	prices := []float64{10.99, 8.99}
// 	fmt.Println(prices[1])

// 	prices = append(prices, 5.99) // Overwrite for a new slice with the appended value
// 	fmt.Println(prices)

// 	// Umpacking List Values
// 	discoutedPrices := []float64{101.99, 45.99, 20.59}
// 	prices = append(prices, discoutedPrices...)
// 	fmt.Println(prices)

// }
