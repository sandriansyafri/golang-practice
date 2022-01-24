package main

import "fmt"

type User struct {
	Name, Username, Email string
}

func (user *User) updateUserName() {
	user.Name = "Sandrian"
}

func main() {
	user1 := User{"A", "A", "A"}
	user1.updateUserName()

	fmt.Println(user1)

}
