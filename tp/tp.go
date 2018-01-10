package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func concatenate(data1 string, data2 string, data3 string) (sum string, err error){
	var concatenateMe string
	concatenateMe = os.Args[4]
	var concatMe string
	var bytedConcat []byte
	var stringifiedData2 string
	var stringifiedData3 string
	data3 = os.Args[3]
	content2, err := ioutil.ReadFile(data2)
	content3, err := ioutil.ReadFile(data3)
	stringifiedData2 = string(content2)
	stringifiedData3 = string(content3)
	concatMe = data1 + stringifiedData2 + stringifiedData3
	bytedConcat = []byte(concatMe)
	err = ioutil.WriteFile(concatenateMe, bytedConcat, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}


func main() {
	var name string
	name = os.Args[1]
	content, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	var contentStr string
	contentStr = string(content)
	fmt.Println(contentStr)

	var writeInto string
	writeInto = os.Args[2]
	ioutil.WriteFile(writeInto, content, 0777)

	var data3 string
	data3 = os.Args[3]
	content3, err := ioutil.ReadFile(data3)
	if err != nil {
		fmt.Println(err)
		return
	}
	var content3Str string
	content3Str = string(content3)
	
	concatenate(name, writeInto, content3Str)
}