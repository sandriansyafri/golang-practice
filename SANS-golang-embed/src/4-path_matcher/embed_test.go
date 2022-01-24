package embed

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed files/*.txt
var filesTxt embed.FS

func TestEmbedMultipleFile(t *testing.T) {
	files, _ := filesTxt.ReadDir("files")
	for _, file := range files {
		if !file.IsDir() {
			fmt.Println("NAME : ", file.Name())
			content, _ := filesTxt.ReadFile("files/" + file.Name())
			fmt.Println(string(content))
			fmt.Println("----------------")
		} else {
			fmt.Println("!OK")
		}

	}
}
