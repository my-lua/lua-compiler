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

// LineStart 单词行开始
func (me *LuaToken) LineStart() int {
	return me.lineStart
}

// LineEnd 单词行结束
func (me *LuaToken) LineEnd() int {
	return me.lineEnd
}

// CharStart 单词列开始
func (me *LuaToken) CharStart() int {
	return me.charStart
}

// CharEnd 单词列结束
func (me *LuaToken) CharEnd() int {
	return me.charEnd
}

// Line 当前行（行开始）
func (me *LuaToken) Line() int {
	return me.LineStart()
}

// Char 当前列（列开始）
func (me *LuaToken) Char() int {
	return me.CharStart()
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
