package lexer

import "regexp"

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
