package lexer

// LuaToken Lua单词
type LuaToken struct {
	tokenType ETokenType
	text      string
	line      int
	column    int
	start     int
	end       int
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

func (me *LuaToken) SetLineStart(value int) {
	me.lineStart = value
}

func (me *LuaToken) SetLineEnd(value int) {
	me.lineEnd = value
}

// NewToken 构造函数
func NewToken(
	tokenType ETokenType,
	text string,
) *LuaToken {
	return &LuaToken{
		tokenType: tokenType,
		text:      text,
	}
}
