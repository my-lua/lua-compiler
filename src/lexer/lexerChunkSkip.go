package lexer

// skipComment 跳过注释
func (me *LuaLexer) skipComment() {

}

// skipInvalidCodes 跳过无效代码
func (me *LuaLexer) skipInvalidCodes() {
	for len(me.chunk) > 0 {
		// if me.chunkStartsWith("--") {
		// 	me.skipComment()
		// }
		if me.chunkStartsWith("\n\r") {
			me.chunkNextN(2)
			me.chunkNewLine()
		} else if me.chunkStartsWith("\r\n") {
			me.chunkNextN(2)
			me.chunkNewLine()
		} else if me.chunkStartsWith("\n") {
			me.chunkNext()
			me.chunkNewLine()
		} else if me.chunkStartsWith("\r") {
			me.chunkNext()
			me.chunkNewLine()
		} else if me.chunkTopIsWhiteSpace() {
			me.chunkNext()
		} else {
			break
		}
	}
}
