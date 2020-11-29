package lexer

import "fmt"

// LuaLexer Lua词法分析器
type LuaLexer struct {
	source     string
	sourceName string
	chunk      string
	curLine    int
	curColumn  int
}

// Source 源代码
func (me *LuaLexer) Source() string {
	return me.source
}

// SourceName 源代码文件名
func (me *LuaLexer) SourceName() string {
	return me.sourceName
}

// CurLine 当前词法解析的行号
func (me *LuaLexer) CurLine() int {
	return me.curLine
}

// CurColumn 当前词法解析的列号
func (me *LuaLexer) CurColumn() int {
	return me.curColumn
}

func isWhiteSpace(c byte) bool {
	switch c {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	}
	return false
}

// NextToken 解析下一个单词
func (me *LuaLexer) NextToken() Token {
	me.skipInvalidCodes()
	if me.chunkIsEmpty() {
		return Token{
			text:      "",
			tokenType: TokenEof,
			line:      me.CurLine(),
			column:    me.CurColumn(),
		}
	}
	fmt.Print(string(me.chunk[0]))
	me.chunkNext()
	return Token{
		text:      "",
		tokenType: TokenIdentifier,
		line:      me.CurLine(),
		column:    me.CurColumn(),
	}
}

// Reset 重置状态机
func (me *LuaLexer) Reset() {
	me.chunk = me.source
	me.curLine = 1
	me.curColumn = 1
}

// Run 运行状态机
func (me *LuaLexer) Run() []Token {
	result := make([]Token, 0)
	for {
		token := me.NextToken()
		result = append(result, token)
		if token.TokenType() == TokenEof {
			break
		}
	}
	return result
}

// NewLexer 构造函数
func NewLexer(source, sourceName string) *LuaLexer {
	return &LuaLexer{
		source:     source,
		sourceName: sourceName,
		chunk:      source,
		curLine:    1,
		curColumn:  1,
	}
}
