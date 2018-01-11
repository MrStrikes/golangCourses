package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	for i := 1; i < len(os.Args); i++{
		args := os.Args[i]
		fmt.Println(args)
		convArgs, _ := strconv.Atoi(args)
		calc := mult(convArgs)
		fmt.Println(calc)
	}
}

func mult(nbr ...int)(answer int){
	for j := 0; j < len(nbr); j++{
		answer = nbr[j] * nbr[j+1]
	}
	return
}