package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var name string
	name = os.Args[1]
	content, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	var contentStr string
	contentStr = string(content)
	fmt.Println(contentStr)

	var byteName []byte
	byteName = []byte(name)
	var writeInto string
	writeInto = os.Args[2]
	fmt.Println(writeInto)
	err = ioutil.WriteFile(writeInto, content, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}
