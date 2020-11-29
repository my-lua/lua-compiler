package lexer

import "strings"

// chunkNext 下一个字符
func (me *LuaLexer) chunkNext() {
	me.chunk = me.chunk[1:]
	me.curColumn++
}

// chunkNext 下n个字符
func (me *LuaLexer) chunkNextN(n int) {
	for i := 0; i < n; i++ {
		me.chunkNext()
	}
}

// chunkNewLine 新一行
func (me *LuaLexer) chunkNewLine() {
	me.curLine++
	me.curColumn = 1
}

// chunkNewLines 新n行
func (me *LuaLexer) chunkNewLineN(n int) {
	for i := 0; i < n; i++ {
		me.chunkNewLine()
	}
}

// chunkStartsWith 检查chunk是否以特定文本开始
func (me *LuaLexer) chunkStartsWith(text string) bool {
	return strings.HasPrefix(me.chunk, text)
}

// chunkTopIsWhiteSpace 检查chunk顶部是否是空白字符
func (me *LuaLexer) chunkTopIsWhiteSpace() bool {
	c := me.chunk[0]
	switch c {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	}
	return false
}

// chunkIsEmpty 检查chunk是否已为空
func (me *LuaLexer) chunkIsEmpty() bool {
	return len(me.chunk) < 1
}
