package embed

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed image.jpg
var fileImage []byte

func TestEmbedString(t *testing.T) {
	result := ioutil.WriteFile("new_image.jpg", fileImage, fs.ModePerm)
	if result != nil {
		panic("ERROR")
	}

}
