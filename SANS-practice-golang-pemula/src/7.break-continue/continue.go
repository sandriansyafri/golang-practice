package main

import "fmt"

func main() {

	for i := 0; i < 9; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)

	}

}
