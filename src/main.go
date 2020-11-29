package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"./lexer"
)

func main() {
	fileName := "test/1.lua"
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("无法从代码文件中读取内容")
	}
	source := string(buf)
	lex := lexer.NewLexer(source, fileName)
	lex.Reset()
	lex.Run()
	exp := regexp.MustCompile(`^\d{4}$`)
	rst := exp.MatchString("1234")
	fmt.Println()
	fmt.Println(rst)
}
