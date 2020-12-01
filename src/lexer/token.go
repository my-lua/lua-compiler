package lexer

import "fmt"

// LuaToken Lua单词
type LuaToken struct {
	tokenType ETokenType
	text      string
	lineStart int
	lineEnd   int
	charStart int
	charEnd   int
}

// TokenType 单词类型
func (me *LuaToken) TokenType() ETokenType {
	return me.tokenType
}

// Text 单词文本
func (me *LuaToken) Text() string {
	return me.text
}

// LineStart s
func (me *LuaToken) LineStart() int {
	return me.lineStart
}

// LineEnd s
func (me *LuaToken) LineEnd() int {
	return me.lineEnd
}

// CharStart s
func (me *LuaToken) CharStart() int {
	return me.charStart
}

// CharEnd s
func (me *LuaToken) CharEnd() int {
	return me.charEnd
}

// Line s
func (me *LuaToken) Line() int {
	return me.LineStart()
}

// Char s
func (me *LuaToken) Char() int {
	return me.CharStart()
}

// Keyword 获取Token内的Lua关键字
func (me *LuaToken) Keyword() ELuaKeyword {
	return NewLuaKeyword(me.Text())
}

// Print 打印Token信息
func (me *LuaToken) Print() {
	fmt.Printf("%s\t%s\n", me.Text(), me.TokenType().Name())
}

// NewLuaToken 构造函数
func NewLuaToken(
	tokenType ETokenType,
	text string,
) *LuaToken {
	return &LuaToken{
		tokenType: tokenType,
		text:      text,
	}
}
