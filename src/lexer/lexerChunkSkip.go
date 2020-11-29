package lexer

// skipComment 跳过注释
func (me *LuaLexer) skipComment() {
	me.chunkNextN(2)
}

// skipInvalidCodes 跳过无效代码
func (me *LuaLexer) skipInvalidCodes() {
	for len(me.chunk) > 0 {
		if me.chunkStartsWith("--") {
			me.skipComment()
		} else if me.chunkTopIsPairOfNewLine() {
			me.chunkNextN(2)
			me.chunkNewLine()
		} else if me.chunkTopIsNewLine() {
			me.chunkNext()
			me.chunkNewLine()
		} else if me.chunkTopIsWhiteSpace() {
			me.chunkNext()
		} else {
			break
		}
	}
}
