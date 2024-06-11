package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Instantiating Structs
	// appUser := User{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		panic(err)
	}

	admin := user.NewAdmin("test@example.com", "123456")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// Alternative Instance
	// appUser := User{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	// ... do something awesome with that gathered data!

	// outputUserDetails(&appUser)

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

// Pass struct value as argument & Using pointers
// func outputUserDetails(u *User) {
// 	...
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
