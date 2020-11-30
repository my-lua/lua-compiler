package lexer

// LuaToken Lua单词
type LuaToken struct {
	tokenType ETokenType
	text      string
	line      int
	column    int
}

// Text 单词文本
func (me *LuaToken) Text() string {
	return me.text
}

// TokenType 单词类型
func (me *LuaToken) TokenType() ETokenType {
	return me.tokenType
}

// Line 单词所在代码行号
func (me *LuaToken) Line() int {
	return me.line
}

// Column 单词所在行内的列位置
func (me *LuaToken) Column() int {
	return me.column
}

// NewToken 构造函数
func NewToken(
	tokenType ETokenType,
	text string,
	line, column int,
) *LuaToken {
	return &LuaToken{
		tokenType: tokenType,
		text:      text,
		line:      line,
		column:    column,
	}
}
