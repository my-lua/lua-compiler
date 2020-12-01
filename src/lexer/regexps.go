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

var _reIdentifierStart = regexp.MustCompile(`^[_a-zA-Z]`)

// ReIdentifierStart s
func (me *LuaLexer) ReIdentifierStart() *regexp.Regexp {
	return _reIdentifierStart
}

var _reNumberStart = regexp.MustCompile(`^[.\d]`)

// ReNumberStart s
func (me *LuaLexer) ReNumberStart() *regexp.Regexp {
	return _reNumberStart
}

var _reNumber = regexp.MustCompile(`^[.\d]\d*`)

// ReNumber s
func (me *LuaLexer) ReNumber() *regexp.Regexp {
	return _reNumber
}

var _reIdentifier = regexp.MustCompile(`^[_a-zA-Z][_a-zA-Z\d]*`)

// ReIdentifier s
func (me *LuaLexer) ReIdentifier() *regexp.Regexp {
	return _reIdentifier
}
