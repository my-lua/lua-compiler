package main

import (
	"fmt"
	"io/ioutil"

	"./lexer"
)

func main() {
	fileName := "test/1.lua"
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("无法从代码文件中读取内容")
	}
	source := string(buf)
	fmt.Println(source)
	lex := lexer.NewLexer(source, fileName)
	lex.Reset()
	lex.Run()
}
