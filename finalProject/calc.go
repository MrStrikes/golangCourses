package main

import (
    "fmt"
    "os"
    "strconv"
)

func mult(nums ...int)  {
    result := 0
    total := 1
    for _, num := range nums {
        result = total * num
        total = result
    }
    fmt.Println(result)
}

func add(nums ...int){
	result := 0
    total := 1
    for _, num := range nums {
        result = total + num
        total = result
    }
    fmt.Println(result)
}

func sub(nums ...int){
	result := 0
    total := 1
    for _, num := range nums {
        result = total - num
        total = result
    }
    fmt.Println(result)
}

func div(nums ...int){
	result := 0
    total := 1
    for _, num := range nums {
        result = total / num
        total = result
    }
    fmt.Println(result)
}

func main() {
    var err error
    var d [] int
    var args= os.Args[1:]
    nums := make([]int, len(args))
    for i := 0; i < len(args); i++ {
        nums[i], err = strconv.Atoi(args[i]);
        if err != nil {
            panic(err)
        }
        strconv.Atoi(args[i])
        d = append(d, nums[i])
    }
    num := d
    if os.Args[1] == "*"{
		mult(num...)
	} else if os.Args[1] == "+"{
		add(num...)
	} else if os.Args[1] == "-"{
		sub(num...)
	} else {
		div(num...)
    }
}