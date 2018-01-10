package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file string
	file = os.Args[1]
	readLine(file)
}

func readLine(path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
