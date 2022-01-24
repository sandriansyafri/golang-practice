package main

import "fmt"

func endApp() {
	fmt.Println("End")
}

func runApp(value int) {
	defer endApp()
	fmt.Println("Run")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApp(0)
}
