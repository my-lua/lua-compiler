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

// Line 单词行范围
func (me *LuaToken) Line() (int, int) {
	return me.LineStart(), me.LineEnd()
}

// CharStart 单词字符开始
func (me *LuaToken) CharStart() int {
	return me.charStart
}

// CharEnd 单词字符结束
func (me *LuaToken) CharEnd() int {
	return me.charEnd
}

// Char 单词字符范围
func (me *LuaToken) Char() (int, int) {
	return me.CharStart(), me.CharEnd()
}

// SetLineStart 设置单词所在行开始行号
func (me *LuaToken) SetLineStart(value int) {
	me.lineStart = value
}

// SetLineEnd 设置单词所在行结束行号
func (me *LuaToken) SetLineEnd(value int) {
	me.lineEnd = value
}

// SetLine 设置单词所在的行范围
func (me *LuaToken) SetLine(start, end int) {
	me.SetLineStart(start)
	me.SetLineEnd(end)
}

// SetCharStart 设置单词所在的字符开始位置
func (me *LuaToken) SetCharStart(value int) {
	me.charStart = value
}

// SetCharEnd 设置单词所在的字符结束位置
func (me *LuaToken) SetCharEnd(value int) {
	me.charEnd = value
}

// SetChar 设置单词所在的字符范围
func (me *LuaToken) SetChar(start, end int) {
	me.SetCharStart(start)
	me.SetCharEnd(end)
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

// Print 打印Token信息（用于调试）
func (me *LuaToken) Print() {
	line1, line2 := me.Line()
	char1, char2 := me.Char()
	fmt.Printf(
		"%s\t%s\t%d %d\t%d %d\n",
		me.Text(),
		me.TokenType().Name(),
		line1,
		line2,
		char1,
		char2,
	)
}
