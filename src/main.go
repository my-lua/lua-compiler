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
明月几时有，
把酒问青天
nihao]==]dongwu`, fileName)
	lex.Reset()
	lex.Run()
	lex.PrintStatus()
	// lex.Run()
}
