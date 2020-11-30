package lexer

import "strings"

// chunkNext 下一个字符
func (me *LuaLexer) chunkNext() string {
	result := me.chunk[:1]
	me.chunk = me.chunk[1:]
	me.curColumn++
	return result
}

// chunkNext 下n个字符
func (me *LuaLexer) chunkNextN(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += me.chunkNext()
	}
	return result
}

// chunkNextSkipN 跳过n个字符
func (me *LuaLexer) chunkNextSkipN(n int) {
	me.chunk = me.chunk[n:]
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

// chunkTopIsNewLine 检查chunk顶部是否是可产生新行的符号
func (me *LuaLexer) chunkTopIsNewLine() bool {
	c := me.chunk[0]
	switch c {
	case '\n', '\r':
		return true
	}
	return false
}

// chunkTopIsPairOfNewLine 检查chunk顶部是否是可产生新行的成对符号
func (me *LuaLexer) chunkTopIsPairOfNewLine() bool {
	return me.chunkStartsWith("\n\r") || me.chunkStartsWith("\r\n")
}

// chunkIsEmpty 检查chunk是否已为空
func (me *LuaLexer) chunkIsEmpty() bool {
	return len(me.chunk) < 1
}
