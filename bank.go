package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

// Write file
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

// Read file
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to read file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}

func main() {

	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		// return
		// panic("Can't continue. Sorry!")
	}

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank :)")
			return
		}

	}

}
