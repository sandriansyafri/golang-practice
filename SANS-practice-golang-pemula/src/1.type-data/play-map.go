package main

import "fmt"

func main() {
	// TODO 1
	var firstMap = make(map[string]string)
	firstMap["NAME"] = "Sandrian"
	fmt.Println(firstMap)

	// TODO 2
	secondMap := map[string]string{
		"name": "Sandrian",
	}

	fmt.Println(secondMap)
}
