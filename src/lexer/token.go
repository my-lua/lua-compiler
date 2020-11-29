package lexer

// Token Lua单词
type Token struct {
	text      string
	tokenType ETokenType
	line      int
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
