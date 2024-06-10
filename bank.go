package main

import "fmt"

func main() {

	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmmount float64
			fmt.Scan(&depositAmmount)
			if depositAmmount <= 0 {
				fmt.Println("Invalid ammount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 3:
			fmt.Print("Withdraw ammount: ")
			var withdrawAmmount float64
			fmt.Scan(&withdrawAmmount)
			if withdrawAmmount <= 0 {
				fmt.Println("Invalid ammount. Must be greater than 0.")
				continue
			}
			if withdrawAmmount > accountBalance {
				fmt.Println("Invalid ammount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank :)")
			return
		}

	}

}
