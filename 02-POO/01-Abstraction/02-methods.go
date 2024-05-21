package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func (u User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func Methods() {
	user := User{"John", "Doe"}
	fmt.Println(user.FullName())
}
