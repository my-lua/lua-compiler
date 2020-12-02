package lexer

// SkipComment 跳过注释
func (me *Chunk) SkipComment() {
	// 跳过注释头（--）
	me.NextN(2)
	// 如果是长注释
	if me.Top().LikeLongString() {
		me.ScanLongString()
	} else {
		for !me.IsEmpty() && !me.Top().IsNewLine() {
			me.Next()
		}
	}
}

// SkipInvalidCodes 跳过无效代码
func (me *Chunk) SkipInvalidCodes() {
	for !me.IsEmpty() {
		if me.Top().LikeComment() {
			me.SkipComment()
		} else if me.Top().IsPairOfNewLine() {
			me.NextN(2)
			me.NewLine()
		} else if me.Top().IsNewLine() {
			me.Next()
			me.NewLine()
		} else if me.Top().IsWhiteSpace() {
			me.Next()
		} else {
			break
		}
	}
}
