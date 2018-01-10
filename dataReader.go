package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var read string
	read = os.Args[1]
	content, err := ioutil.ReadFile(read)
	if err != nil {
		fmt.Println(err)
		return
	}
	var contentStr string
	contentStr = string(content)
	fmt.Println(contentStr)
}
