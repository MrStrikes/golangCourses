package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	var a string
	var b string
	var c int
	a = os.Args[1]
	b = os.Args[2]
	mult, _ := strconv.Atoi(a)
	mult2, _ := strconv.Atoi(b)
	c = mult * mult2
	fmt.Println(c)
}
