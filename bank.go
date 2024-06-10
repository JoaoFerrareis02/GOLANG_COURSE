package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		// return
		// panic("Can't continue. Sorry!")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank :)")
			return
		}

	}

}
