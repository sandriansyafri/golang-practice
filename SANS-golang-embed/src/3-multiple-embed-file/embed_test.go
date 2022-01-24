package embed

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestEmbedMultipleFile(t *testing.T) {
	readFile, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(readFile))

}
