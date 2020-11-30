package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// fileName := "test/1.lua"
	// buf, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	panic("无法从代码文件中读取内容")
	// }
	// source := string(buf)
	// lex := lexer.NewLexer(source, fileName)
	// lex.Reset()
	// lex.Run()
	// exp := regexp.MustCompile(`^\d{4}$`)
	// rst := exp.MatchString("1234")
	// fmt.Println()
	// fmt.Println(rst)

	scanLongString(`[====[
古朗月行

作者：李白

小时不识月，呼作白玉盘。
又疑瑶台镜，飞在青云端。
]====]  	`)
}

func scanLongString(text string) string {
	reNewLine := regexp.MustCompile("\r\n|\n\r|\n|\r")
	reLongStringOpeningBracket := regexp.MustCompile(`^\[(=*)\[`)
	matchs := reLongStringOpeningBracket.FindStringSubmatch(text)
	if len(matchs) < 2 {
		panic("长字符串头部括号不匹配")
	}
	// 头部长度
	openingBracketLen := len(matchs[0])
	// 等号个数
	equalSignNum := len(matchs[1])
	longStringClosingBracket := `\]={` +
		strconv.FormatInt(int64(equalSignNum), 10) +
		`}\]`
	reLongStringClosingBracket := regexp.MustCompile(longStringClosingBracket)
	text = text[openingBracketLen:]
	indexs := reLongStringClosingBracket.FindStringIndex(text)
	if len(indexs) < 2 {
		panic("长字符串没有正确的尾部括号")
	}
	closingBracketStart := indexs[0]
	text = text[:closingBracketStart]
	// 从代码文本中采集的换行符统一替换为\n
	text = reNewLine.ReplaceAllString(text, "\n")
	// 如果行首为换行符，则删除
	if strings.HasPrefix(text, "\n") {
		text = text[1:]
	}
	fmt.Println(text)
	return ""
}
