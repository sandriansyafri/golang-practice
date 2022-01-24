package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklistName Blacklist) {
	if blacklistName(name) {
		fmt.Println("blocked!")
	}
	fmt.Println("Hello", name)
}

func main() {
	registerUser("admin", func(name string) bool {
		return name == "admin"
	})
}
