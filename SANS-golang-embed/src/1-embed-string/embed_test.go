package embed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed test.txt
var version1 string

//go:embed test.txt
var version2 string

func TestEmbedString(t *testing.T) {
	fmt.Println(version1)
	fmt.Println(version2)
}
