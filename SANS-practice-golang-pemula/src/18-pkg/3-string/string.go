package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello Rian", "Rian")) // true
	fmt.Println(strings.ToLower("AAAAAAAAAAAAAAA"))     // aaaaaaaaaa
	fmt.Println(strings.ToUpper("BBBBBBBBBBBBBBBB"))    // bbbbbbbbbb
	fmt.Println(strings.Compare("RIAN", "RIAN"))        // 0
	fmt.Println(strings.Compare("RIAN", "rian"))        // -1
	fmt.Println(strings.Trim("            RIAN", " "))  // Rian
	fmt.Println(strings.ReplaceAll("Hello eeeeeee", "ello", "i"))
}
