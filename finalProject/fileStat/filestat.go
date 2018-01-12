package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The file is", fi.Size(), "bytes")
	inFile, _ := ioutil.ReadFile("file.txt")
	for i := 0; i < len(inFile); i++ {
		runeFile := []rune(string(inFile[i]))
		fmt.Println(runeFile)
	}
	fmt.Println(maping(inFile))
}

func maping(byteContent []byte) map[rune]int {
	runeContent := []rune(string(byteContent))
	myMap := make(map[rune]int)
	for _, j := range runeContent {
		myMap[j] = myMap[j] + 1
	}
	fmt.Println(myMap)
	return myMap
}
