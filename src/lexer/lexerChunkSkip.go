package lexer

// skipComment 跳过注释
func (me *LuaLexer) skipComment() {
	// 跳过注释头（--）
	me.chunkNextN(2)
	// 如果是长注释
	if me.ReLongStringOpeningBracket().FindString(me.chunk) != "" {
		me.scanLongString()
	} else {
		for !me.chunkIsEmpty() && !me.chunkTopIsNewLine() {
			me.chunkNext()
		}
	}
}

// skipInvalidCodes 跳过无效代码
func (me *LuaLexer) skipInvalidCodes() {
	for !me.chunkIsEmpty() {
		if me.chunkTopIsComment() {
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
