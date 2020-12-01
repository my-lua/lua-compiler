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
	// lex.PrintStatus()
	// lex.Run()

	smap := map[string]int{}
	fmt.Println(smap["xxss"])
	re := regexp.MustCompile(`^[_a-zA-Z][_a-zA-Z\d]*`)
	fmt.Println(re.MatchString("_ass我K1293abc"))
}
