package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	inFile, _ := os.Open(os.Args[1])
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	fmt.Printf("HTML code of %s \n", scanner.Text())
	for scanner.Scan() {
		wg.Add(1)
		go multi(scanner.Text())
	}
	wg.Wait()
}

func multi(text string) {
	resp, err := http.Get(text)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var aroNum int
	aroNum = bytes.Count([]byte(html), []byte("@"))
	fmt.Printf("%s\n", html)
	fmt.Printf("The number of @ is %d", aroNum)
	wg.Done()
}
