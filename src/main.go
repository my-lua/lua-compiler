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

	re := regexp.MustCompile(`^('.*?')|(".*?")`)
	str := `"123"4" +  "sdf"`
	fmt.Println(re.MatchString(str), re.FindString(str))

	s := "abcd"
	fmt.Println(string([]rune(s)[2:]))
}
