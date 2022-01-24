package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	var x map[string]string

	if x == nil {
		fmt.Println("(empty)")
	} else {
		fmt.Println(x)
	}

}
