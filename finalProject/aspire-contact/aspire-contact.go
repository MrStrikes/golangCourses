package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	inFile, _ := os.Open(os.Args[1])
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	fmt.Printf("HTML code of %s \n", scanner.Text())
	for scanner.Scan() {
		resp, err := http.Get(scanner.Text())
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		html, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", html)
	}
}
