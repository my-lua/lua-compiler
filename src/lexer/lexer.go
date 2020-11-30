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

// Reset 重置状态机
func (me *LuaLexer) Reset() {
	me.chunk = me.source
	me.curLine = 1
	me.curColumn = 1
}

// Run 运行状态机
func (me *LuaLexer) Run() []LuaToken {
	result := make([]LuaToken, 0)
	for {
		token := me.NextToken()
		result = append(result, token)
		if token.TokenType() == TokenEof {
			break
		}
	}
	return result
}

// PrintStatus s
func (me *LuaLexer) PrintStatus() {
	fmt.Printf("file: %s\n", me.SourceName())
	fmt.Printf("line: %d\n", me.CurLine())
	fmt.Printf("column: %d\n", me.CurColumn())
	if len(me.chunk) < 10 {
		fmt.Printf("chunk: %s...\n", me.chunk[:len(me.chunk)])
	} else {
		fmt.Printf("chunk: %s...\n", me.chunk[:10])
	}
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
