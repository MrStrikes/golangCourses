package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

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
	err = ioutil.WriteFile(writeInto, content, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	var concatenateMe string
	concatenateMe = os.Args[4]
	var concatMe string
	var data3 string
	var bytedConcat []byte
	var stringifiedData2 string
	var stringifiedData3 string
	data3 = os.Args[3]
	content2, err := ioutil.ReadFile(writeInto)
	content3, err := ioutil.ReadFile(data3)
	stringifiedData2 = string(content2)
	stringifiedData3 = string(content3)
	concatMe = contentStr + stringifiedData2 + stringifiedData3
	bytedConcat = []byte(concatMe)
	err = ioutil.WriteFile(concatenateMe, bytedConcat, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func concatenate(data1 string, data2 string, data3 string) (sum string, err error){
	return
}