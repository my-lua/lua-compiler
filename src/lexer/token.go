package lexer

// Token Lua单词
type Token struct {
	tokenType ETokenType
	text      string
	line      int
	column    int
}

// Text 单词文本
func (me *Token) Text() string {
	return me.text
}

// TokenType 单词类型
func (me *Token) TokenType() ETokenType {
	return me.tokenType
}

// Line 单词所在代码行号
func (me *Token) Line() int {
	return me.line
}

// Column 单词所在行内的列位置
func (me *Token) Column() int {
	return me.column
}
