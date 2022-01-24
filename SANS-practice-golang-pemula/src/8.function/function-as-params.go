package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, spamFilter Filter) string {
	return "Hello " + spamFilter(name)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	fmt.Println(sayHelloWithFilter("Anjing", spamFilter))
}
