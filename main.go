package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/* var content []byte
	var err error */ //Other way than :=

	content, err := ioutil.ReadFile("text.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(content)

	var contentStr string
	contentStr = string(content)
	fmt.Println(contentStr)
}
