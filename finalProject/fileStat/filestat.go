package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The file is", fi.Size(), "bytes")
	inFile, _ := ioutil.ReadFile("file.txt")
	res := maping(inFile)
	count(res)
}

func maping(byteContent []byte) map[rune]int {
	runeContent := []rune(string(byteContent))
	myMap := make(map[rune]int)
	for _, j := range runeContent {
		myMap[j] = myMap[j] + 1
	}
	return myMap
}

func count(Map map[rune]int) {
	var maxCount int
	var maxLetter []rune
	minCount := 99999999
	var minLetter []rune
	for i, j := range Map {
		if unicode.IsLetter(i) {
			if j > maxCount {
				maxCount = j
				maxLetter = []rune{i}
			} else if j >= maxCount {
				maxCount = j
				maxLetter = append(maxLetter, i)
			}
			if j < minCount {
				minCount = j
				minLetter = []rune{i}
			} else if j <= minCount {
				minCount = j
				minLetter = append(minLetter, i)
			}
		}
	}
	if len(minLetter) < 2 {
		fmt.Printf("Least recurrent letter is %s with %d occurrences\n", strings.Split(string(minLetter), ""), minCount)
	} else {
		fmt.Printf("Least recurrent letters are %s with %d occurrences\n", strings.Split(string(minLetter), ""), minCount)
	}

	if len(maxLetter) < 2 {
		fmt.Printf("Most recurrent letter is %s with %d occurrences\n", strings.Join(strings.Split(string(maxLetter), ""), ","), maxCount)
	} else {
		fmt.Printf("Most recurrent letters are %q with %d occurrences\n", strings.Join(strings.Split(string(maxLetter), ""), ", "), maxCount)
	}

}
