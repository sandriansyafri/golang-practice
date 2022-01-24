package main

import "fmt"

func endApp() {
	fmt.Println("End.")
}

func runApp(error bool) {
	if error {
		panic("ERROR!")
	} else {
		fmt.Println("ok")
	}
}
func main() {
	runApp(true)
}
