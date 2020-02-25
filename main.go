package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	
	start := time.Now()
	//Root directory path, where version values need to be injected
	root := flag.String("dir", "", "Root Directory path where HTML files are stored.")
	//Optional: version can be passed from command line, else it should take a random number within 999.
	version := flag.Int("ver", rand.Intn(999), "Version number which whill be used to update the JS or CSS files version. If not specified random number is generated within 999")
	flag.Parse()
	if *root == "" {
		fmt.Println("Directory name is mandtory. Please check --help")
		os.Exit(2)
	}

	vFiles, err := getHTMLFiles(*root)

	if err != nil {
		fmt.Println("Error in reading files, please check path", err)
		os.Exit(2)
	}

	err = readFileContents(vFiles)

	if err != nil {
		fmt.Println("Error in Reading file contents", err)
		os.Exit(2)
	}

	for _, vFile := range vFiles {
		vFile.updateVersion(*version)
		vFile.writeToFile()
	}

	fmt.Println("Total time taken : ", time.Since(start))
}
