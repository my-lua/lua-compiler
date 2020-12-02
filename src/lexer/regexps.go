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

var _reLongStringBracketStart = regexp.MustCompile(`^\[(=*)\[`)

// ReLongStringBracketStart s
func (me *LuaLexer) ReLongStringBracketStart() *regexp.Regexp {
	return _reLongStringBracketStart
}

// ReLongStringBracketEnd s
func (me *LuaLexer) ReLongStringBracketEnd(n int) *regexp.Regexp {
	// 构建尾部括号正则表达式
	longStringBracketEnd := `\]={` +
		strconv.FormatInt(int64(n), 10) +
		`}\]`
	return regexp.MustCompile(longStringBracketEnd)
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

// var _reShortStr = regexp.MustCompile(`
// 	(?s)
// 	(^'(\\\\|\\'|\\\n|\\z\s*|[^'\n])*') |
// 	(^"(\\\\|\\"|\\\n|\\z\s*|[^"\n])*")
// `)

var _reShortStr = regexp.MustCompile(`^('.*?')|(".*?")`)

// ReShortStr 短字符串（目前简单写的有错误）
func (me *LuaLexer) ReShortStr() *regexp.Regexp {
	return _reShortStr
}
