package lexer

import "fmt"

// LuaLexer Lua词法分析器
type LuaLexer struct {
	source     string
	sourceName string
	chunk      *Chunk
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
	return me.chunk.CurLine()
}

// CurChar 当前词法解析的列号
func (me *LuaLexer) CurChar() int {
	return me.chunk.CurChar()
}

// Reset 重置状态机
func (me *LuaLexer) Reset() {
	me.chunk = NewChunk(me.source)
}

// Run 运行状态机
func (me *LuaLexer) Run() []*LuaToken {
	result := make([]*LuaToken, 0)
	for {
		token := me.NextToken()
		token.Print()
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
	fmt.Printf("column: %d\n", me.CurChar())
	if len(me.chunk.Text()) < 10 {
		fmt.Printf("chunk: %s...\n", me.chunk.Text()[:len(me.chunk.Text())])
	} else {
		fmt.Printf("chunk: %s...\n", me.chunk.Text()[:10])
	}
}

// NewLexer 构造函数
func NewLexer(source, sourceName string) *LuaLexer {
	return &LuaLexer{
		source:     source,
		sourceName: sourceName,
		chunk:      NewChunk(source),
	}
}
