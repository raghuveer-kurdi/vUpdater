package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

	re := regexp.MustCompile(`.js(\?*.*\d+)`)
	for _, file := range files {
		fileContents, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file contents")
			os.Exit(2)
		}
		lines := strings.Split(string(fileContents), "\n")
		for _, line := range lines {
			if strings.Contains(line, "script") {
				fmt.Println(re.ReplaceAllString(line, ".js?v=123"))
			}
		}
	}
}
