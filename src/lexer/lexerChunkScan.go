package lexer

import (
	"regexp"
	"strconv"
	"strings"
)

func (me *LuaLexer) scan(re *regexp.Regexp) string {
	return ""
}

func (me *LuaLexer) scanShortString() string {
	return ""
}

// scanLongString 扫描长字符串
func (me *LuaLexer) scanLongString() string {
	// 检查长字符串头部括号
	matchs := me.ReLongStringOpeningBracket().FindStringSubmatch(me.chunk)
	if len(matchs) < 2 {
		panic("lexer scanLongString: 长字符串头部括号不匹配")
	}
	// 获取头部括号长度, 其中等号个数
	openingBracketLen, equalSignNum := len(matchs[0]), len(matchs[1])
	// 跳过头部括号
	me.chunkNextN(openingBracketLen)
	// 构建尾部括号正则表达式
	longStringClosingBracket := `\]={` +
		strconv.FormatInt(int64(equalSignNum), 10) +
		`}\]`
	reLongStringClosingBracket := regexp.MustCompile(longStringClosingBracket)
	// 向后寻找正确的尾部括号
	indexs := reLongStringClosingBracket.FindStringIndex(me.chunk)
	if len(indexs) < 2 {
		panic("lexer scanLongString: 长字符串没有正确的尾部括号")
	}
	// 获取尾部括号开始，结束位置
	closingBracketStartIndex := indexs[0]
	// 获取并处理字符串
	result := ""
	for i := 0; i < closingBracketStartIndex; i++ {
		if me.chunkTopIsPairOfNewLine() {
			me.chunkNextN(2)
			me.chunkNewLine()
			result += "\n"
		} else if me.chunkTopIsNewLine() {
			me.chunkNext()
			me.chunkNewLine()
			result += "\n"
		} else {
			result += me.chunkNext()
		}
	}
	me.chunkNextN(openingBracketLen)
	// 如果行首为换行符，则删除
	if strings.HasPrefix(result, "\n") {
		result = result[1:]
	}
	return result
}

func (me *LuaLexer) scanNumber() string {
	return ""
}

func (me *LuaLexer) scanIdentifier() string {
	return ""
}
