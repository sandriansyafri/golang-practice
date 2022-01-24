package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println(message)
	fmt.Println("End.")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR!")
	} else {
		fmt.Println("ok")
	}
}
func main() {
	runApp(true)
}
