package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "test/1.lua"
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("无法从代码文件中读取内容")
	}
	fmt.Println(string(buf))
}
