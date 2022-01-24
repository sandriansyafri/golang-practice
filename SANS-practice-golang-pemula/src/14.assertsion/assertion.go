package main

import (
	"fmt"
)

func test() interface{} {
	return 1
}
func main() {
	var result = test()

	switch value := result.(type) {
	case string:
		fmt.Println("value :", value, "is String")
	case int:
		fmt.Println("value :", value, "is Int")
	default:
		fmt.Println("Unknown")
	}

}
