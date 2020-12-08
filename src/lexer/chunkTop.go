package lexer

import (
	"regexp"
	"strings"
)

// ChunkTop Chunk头部
type ChunkTop string

// Char 顶部字符
func (me ChunkTop) Char() byte {
	return me[0]
}

// CharStr 顶部字符（字符串形式）
func (me ChunkTop) CharStr() string {
	return string(me.Char())
}

// Like 传入正则表达式，进行匹配判断
func (me ChunkTop) Like(re *regexp.Regexp) bool {
	return re.MatchString(string(me))
}

// StartsWith 检查chunk顶部是否以某给定字符串开头
func (me ChunkTop) StartsWith(text string) bool {
	return strings.HasPrefix(string(me), text)
}

// IsWhiteSpace 检查chunk顶部是否是空白字符
func (me ChunkTop) IsWhiteSpace() bool {
	switch me.CharStr() {
	case "\t", "\n", "\v", "\f", "\r", " ":
		return true
	}
	return false
}

// IsNewLine 检查chunk顶部是否是可产生新行的符号
func (me ChunkTop) IsNewLine() bool {
	switch me.CharStr() {
	case "\n", "\r":
		return true
	}
	return false
}

// IsPairOfNewLine 检查chunk顶部是否是可产生新行的成对符号
func (me ChunkTop) IsPairOfNewLine() bool {
	return me.StartsWith("\r\n") || me.StartsWith("\n\r")
}

// LikeNumber 像是数字
func (me ChunkTop) LikeNumber() bool {
	return me.Like(ReNumberStart())
}

// LikeShortString 像是短字符串
func (me ChunkTop) LikeShortString() bool {
	return me.StartsWith("'") || me.StartsWith("\"")
}

// LikeLongString 像是长字符串
func (me ChunkTop) LikeLongString() bool {
	return me.Like(ReLongStringBracketStart())
}

// LikeIdentifier 像是标识符
func (me ChunkTop) LikeIdentifier() bool {
	return me.Like(ReIdentifierStart())
}

// LikeComment 像是注释
func (me ChunkTop) LikeComment() bool {
	return me.StartsWith("--")
}
