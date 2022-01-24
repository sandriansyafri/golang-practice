package main

import "fmt"

func main() {
	person := map[string]string{
		"name":     "Sandrian Syafri",
		"username": "sandriansyafri",
	}
	person["email"] = "sandrian@gmail.com"

	person2 := make(map[string]string)
	person2["name"] = "Example"
	person2["username"] = "example"
	person2["email"] = "example@gmail.com"
	fmt.Println("BEFORE :", person2)
	delete(person2, "email")
	fmt.Println("AFTER", person2)

}
