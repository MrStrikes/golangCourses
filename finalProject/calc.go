package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	for i := 1; i < len(os.Args); i++{
		var calc string
		args := os.Args[i]
		calc, _ = strconv.Atoi(args)
		fmt.Println(calc)
	}
}
