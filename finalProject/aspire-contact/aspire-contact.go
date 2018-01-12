package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	urls, _ := ioutil.ReadFile("URL.txt")
	stringUrls := string(urls)
	fmt.Printf("HTML code of %s ...\n", urls)
	resp, err := http.Get(stringUrls)
	if err != nil {
		panic(err)
	}
}
