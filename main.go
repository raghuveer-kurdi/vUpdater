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

	root := "c:\\Users\\raghuveer.k\\Desktop\\Temp"
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

	jsDQuote := regexp.MustCompile(`\.js(\?*.*\d*")`)
	jsSQuote := regexp.MustCompile(`\.js(\?*.*\d*')`)
	cssDQuote := regexp.MustCompile(`\.css(\?*\w*=*\d*")`)
	cssSQuote := regexp.MustCompile(`\.css(\?*\w*=*\d*')`)
	for _, file := range files {
		fileContents, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file contents")
			os.Exit(2)
		}
		lines := strings.Split(string(fileContents), "\n")
		fmt.Println("--------file name : ", file, "-----------------")
		for _, line := range lines {
			if jsDQuote.MatchString(line) {
				fmt.Println(jsDQuote.ReplaceAllString(line, ".js?v=123\""))
			} else if jsSQuote.MatchString(line) {
				fmt.Println(jsSQuote.ReplaceAllString(line, ".js?v=123'"))
			} else if cssDQuote.MatchString(line) {
				fmt.Println(cssDQuote.ReplaceAllString(line, ".css?v=123\""))
			} else if cssSQuote.MatchString(line) {
				fmt.Println(cssDQuote.ReplaceAllString(line, ".css?v=123'"))
			}
		}
	}
}
