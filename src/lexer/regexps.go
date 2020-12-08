package lexer

import (
	"regexp"
	"strconv"
)

var _reNewLine = regexp.MustCompile("\r\n|\n\r|\n|\r")

// ReNewLine 换行符
func ReNewLine() *regexp.Regexp {
	return _reNewLine
}

var _reLongStringBracketStart = regexp.MustCompile(`^\[(=*)\[`)

// ReLongStringBracketStart 长字符串开始
func ReLongStringBracketStart() *regexp.Regexp {
	return _reLongStringBracketStart
}

// ReLongStringBracketEnd 长字符串结束
func ReLongStringBracketEnd(n int) *regexp.Regexp {
	// 构建尾部括号正则表达式
	longStringBracketEnd := `\]={` +
		strconv.FormatInt(int64(n), 10) +
		`}\]`
	return regexp.MustCompile(longStringBracketEnd)
}

var _reIdentifierStart = regexp.MustCompile(`^[_a-zA-Z]`)

// ReIdentifierStart 标识符开始
func ReIdentifierStart() *regexp.Regexp {
	return _reIdentifierStart
}

var _reNumberStart = regexp.MustCompile(`^[.\d]`)

// ReNumberStart 数字开始
func ReNumberStart() *regexp.Regexp {
	return _reNumberStart
}

var _reNumber = regexp.MustCompile(`^[.\d]\d*`)

// ReNumber 数字
func ReNumber() *regexp.Regexp {
	return _reNumber
}

var _reIdentifier = regexp.MustCompile(`^[_a-zA-Z][_a-zA-Z\d]*`)

// ReIdentifier 标识符
func ReIdentifier() *regexp.Regexp {
	return _reIdentifier
}

var _reShortStr = regexp.MustCompile(`^('.*?')|(".*?")`)

// ReShortStr 短字符串（目前简单写的有错误）
func ReShortStr() *regexp.Regexp {
	return _reShortStr
}
