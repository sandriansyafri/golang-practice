package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed file-embed/version/version.txt
var version string

//go:embed file-embed/txt/*.txt
var txt_files embed.FS

//go:embed file-embed/img/image.jpg
var fileImage []byte

func main() {
	fmt.Println(version)
	entries, _ := txt_files.ReadDir("file-embed/txt")

	for _, entry := range entries {
		content, _ := txt_files.ReadFile("file-embed/txt/" + entry.Name())
		fmt.Println(entry.Name(), ":", string(content))
	}

	err := ioutil.WriteFile("file-embed/img/image-new.jpg", fileImage, fs.FileMode(fs.ModePerm))
	if err != nil {
		panic("FAILED")
	}

}
