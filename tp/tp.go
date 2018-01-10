package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func read(filename string) (content []byte){
	content, _ = ioutil.ReadFile(filename)
	return
}

func write(filename string, content[]byte){
	ioutil.WriteFile(filename, content, 0644)
}
 
func concatenate(datas ...string) (sum []byte, err error){
	for i := 0; i < len(datas); i++{
		f := read(datas[i])
		sum = append(sum, f...)
	}
	return
}

func main() {
	if len(os.Args) < 5{
		fmt.Println("I need at least 4 files!")
		return
	}
	sum, _ := concatenate(os.Args[1], os.Args[2], os.Args[3])
	write(os.Args[4], sum)
}