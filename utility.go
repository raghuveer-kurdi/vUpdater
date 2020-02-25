package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

const (
	jsDRegEx    = `\.js(\?*.*\d*")`
	jsSRegEx    = `\.js(\?*.*\d*')`
	cssDRegEx   = `\.css(\?*\w*=*\d*")`
	cssSRegEx   = `\.css(\?*\w*=*\d*')`
	jsDVFormat  = ".js?v=%d\""
	jsSVFormat  = ".js?v=%d'"
	cssDVFormat = ".css?v=%d\""
	cssSVFormat = ".css?v=%d'"
)

func modifyLine(line string, version int) string {
	//Utility method which inserts the version.
	jsDQuote := regexp.MustCompile(jsDRegEx)
	jsSQuote := regexp.MustCompile(jsSRegEx)
	cssDQuote := regexp.MustCompile(cssDRegEx)
	cssSQuote := regexp.MustCompile(cssSRegEx)

	if jsDQuote.MatchString(line) {
		line = jsDQuote.ReplaceAllString(line, fmt.Sprintf(jsDVFormat, version))
	} else if jsSQuote.MatchString(line) {
		line = jsSQuote.ReplaceAllString(line, fmt.Sprintf(jsSVFormat, version))
	} else if cssDQuote.MatchString(line) {
		line = cssDQuote.ReplaceAllString(line, fmt.Sprintf(cssDVFormat, version))
	} else if cssSQuote.MatchString(line) {
		line = cssSQuote.ReplaceAllString(line, fmt.Sprintf(cssSVFormat, version))
	}

	return line
}

func getHTMLFiles(rootDir string) ([]vFile, error) {
	vFiles := []vFile{}
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || (filepath.Ext(path) != ".html" && filepath.Ext(path) != ".htm") {
			return nil
		}
		fileDetail := vFile{name: info.Name(), path: path}
		vFiles = append(vFiles, fileDetail)
		return nil
	})

	return vFiles, err
}

func readFileContents(vFiles []vFile) error {
	var err error
	for i := 0; i < len(vFiles); i++ {
		err = vFiles[i].readContent()
		if err != nil {
			break
		}
	}
	return err
}
