package lexer

import (
	"regexp"
	"strconv"
)

var _reNewLine = regexp.MustCompile("\r\n|\n\r|\n|\r")

// ReNewLine s
func (me *LuaLexer) ReNewLine() *regexp.Regexp {
	return _reNewLine
}

var _reLongStringOpeningBracket = regexp.MustCompile(`^\[(=*)\[`)

// ReLongStringOpeningBracket s
func (me *LuaLexer) ReLongStringOpeningBracket() *regexp.Regexp {
	return _reLongStringOpeningBracket
}

// ReLongStringClosingBracket s
func (me *LuaLexer) ReLongStringClosingBracket(n int) *regexp.Regexp {
	// 构建尾部括号正则表达式
	longStringClosingBracket := `\]={` +
		strconv.FormatInt(int64(n), 10) +
		`}\]`
	return regexp.MustCompile(longStringClosingBracket)
}
