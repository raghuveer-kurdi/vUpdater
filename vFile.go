package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type vFile struct {
	name     string
	path     string
	content  string
	uContent string
}

//this file has file structure and other utility methods like reading and writing to files. 

func (vf vFile) print(modified bool) {
	fmt.Println("********************** ", vf.name, " ********************")
	fmt.Println("Path : ", vf.path)
	fmt.Println("Content : ")
	if !modified {
		fmt.Println(vf.content)
	} else {
		fmt.Println(vf.uContent)
	}
}

func (vf vFile) println() {}

func (vf *vFile) updateVersion(version int) {
	lines := strings.Split(vf.content, "\n")
	uLines := []string{}
	for _, line := range lines {
		uLines = append(uLines, modifyLine(line, version))
	}
	vf.uContent = strings.Join(uLines, "\n")
}

func (vf *vFile) readContent() error {
	bs, err := ioutil.ReadFile(vf.path)
	if err != nil {
		return err
	}
	content := string(bs)
	vf.content = content
	return nil
}

func (vf vFile) writeToFile() error {
	return ioutil.WriteFile(vf.path, []byte(vf.uContent), 0666)
}
