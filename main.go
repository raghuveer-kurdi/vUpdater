package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var files []string

	root := "c:\\Users\\raghu\\Desktop\\Temp"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || (filepath.Ext(path) != ".html" && filepath.Ext(path) != ".htm") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
