package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO 1
	/**
		write args in terminal
		go run [name.go] [arg1] [arg2] [arg...]
	**/
	// args := os.Args // check location file os.exe
	// fmt.Println(args)

	// TODO 2
	// hostname, errs := os.Hostname()
	// if errs == nil {
	// 	fmt.Println(hostname)
	// } else {
	// 	fmt.Println(errs.Error())
	// }

	// TODO
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

}
