package main

import "fmt"

func (u User) displayEmail() {
	fmt.Printf("User %d Email: %s\n", u.uid, u.email)
}

func main() {
	userExample := User{
		uid:   1,
		name:  "bob",
		email: "bob@example.com",
		phone: "123-456-7890",
	}
	userExample.displayEmail()
}
