package lexer

// LuaLexer Lua词法分析器
type LuaLexer struct {
	chunk     string
	chunkName string
	curLine   int
}

// Chunk 获取源代码
func (me *LuaLexer) Chunk() string {
	return me.chunk
}

// ChunkName 源代码文件名
func (me *LuaLexer) ChunkName() string {
	return me.chunkName
}

// CurLine 当前词法解析的行号
func (me *LuaLexer) CurLine() int {
	return me.curLine
}

// NextToken 解析下一个单词
func (me *LuaLexer) NextToken() {

}

// NewLexer 构造函数
func NewLexer(chunk, chunkName string) *LuaLexer {
	return &LuaLexer{
		chunk:     chunk,
		chunkName: chunkName,
		curLine:   1,
	}
}
