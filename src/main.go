package main

import (
	"./lexer"
)

func main() {
	fileName := "test/1.lua"
	// buf, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	panic("无法从代码文件中读取内容")
	// }
	// source := string(buf)
	lex := lexer.NewLexer(`[==[

	asddf
	]==]   asd 搜索dongwu
	`, fileName)
	lex.Reset()
	lex.Run()
	lex.PrintStatus()
	// lex.Run()
}
