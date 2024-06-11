package user

import (
	"errors"
	"fmt"
	"time"
)

// Defining a struct type
// Export structs
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Struct Embedding
type Admin struct {
	email    string
	password string
	User
}

// Exposing Methods

// Methods - Optional pointer
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// Mutation Methods - Mandatory pointer
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}
}

// Constructor Function (NewUser) - Better use a pointer
// & Using for validation & Diffenrent Constructor Function Name (New)
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
