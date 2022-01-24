package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "localhost")
	username := flag.String("root", "root", "username")
	password := flag.String("password", "root", "password")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*username)
	fmt.Println(*password)
}
