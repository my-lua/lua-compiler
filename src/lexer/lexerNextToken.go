package lexer

// NextToken 解析下一个单词
func (me *LuaLexer) NextToken() *LuaToken {
	// 跳过无效代码
	me.skipInvalidCodes()
	// 如果已经没有要解析的代码，则返回EOF
	if me.chunkIsEmpty() {
		return NewLuaToken(TokenEof, "")
	}
	// 获取顶部字符
	topChar := me.chunkTopCharStr()
	switch topChar {
	case ";":
		me.chunkNext()
		return NewLuaToken(TokenSepSemi, topChar)
	case ",":
		me.chunkNext()
		return NewLuaToken(TokenSepComma, topChar)
	case "(":
		me.chunkNext()
		return NewLuaToken(TokenSepLParen, topChar)
	case ")":
		me.chunkNext()
		return NewLuaToken(TokenSepRParen, topChar)
	case "[":
		if me.chunkTopLikeLongString() {
			return NewLuaToken(TokenString, me.scanLongString())
		}
		me.chunkNext()
		return NewLuaToken(TokenSepLBrack, topChar)
	case "]":
		me.chunkNext()
		return NewLuaToken(TokenSepRBrack, topChar)
	case "{":
		me.chunkNext()
		return NewLuaToken(TokenSepLCurly, topChar)
	case "}":
		me.chunkNext()
		return NewLuaToken(TokenSepRCurly, topChar)
	case "+":
		me.chunkNext()
		return NewLuaToken(TokenOpAdd, topChar)
	case "-":
		me.chunkNext()
		return NewLuaToken(TokenOpSub, topChar)
	case "*":
		me.chunkNext()
		return NewLuaToken(TokenOpMul, topChar)
	case "/":
		me.chunkNext()
		return NewLuaToken(TokenOpDiv, topChar)
	case "%":
		me.chunkNext()
		return NewLuaToken(TokenOpMod, topChar)
	case "^":
		me.chunkNext()
		return NewLuaToken(TokenOpPow, topChar)
	case "&":
		me.chunkNext()
		return NewLuaToken(TokenOpBAnd, topChar)
	case "|":
		me.chunkNext()
		return NewLuaToken(TokenOpBOr, topChar)
	case "#":
		me.chunkNext()
		return NewLuaToken(TokenOpLen, topChar)
	case ":":
		if me.chunkStartsWith("::") {
			me.chunkNextN(2)
			return NewLuaToken(TokenSepLabel, "::")
		}
		me.chunkNext()
		return NewLuaToken(TokenSepColon, topChar)
	case "~":
		if me.chunkStartsWith("~=") {
			me.chunkNextN(2)
			return NewLuaToken(TokenOpNe, "~=")
		}
		me.chunkNext()
		return NewLuaToken(TokenOpWave, topChar)
	case "=":
		if me.chunkStartsWith("==") {
			me.chunkNextN(2)
			return NewLuaToken(TokenOpEq, "==")
		}
		me.chunkNext()
		return NewLuaToken(TokenOpAssign, topChar)
	case "<":
		if me.chunkStartsWith("<<") {
			me.chunkNextN(2)
			return NewLuaToken(TokenOpShl, "<<")
		} else if me.chunkStartsWith("<=") {
			me.chunkNextN(2)
			return NewLuaToken(TokenOpLe, "<=")
		}
		me.chunkNext()
		return NewLuaToken(TokenOpLt, topChar)
	case ">":
		if me.chunkStartsWith(">>") {
			me.chunkNextN(2)
			return NewLuaToken(TokenOpShr, ">>")
		} else if me.chunkStartsWith(">=") {
			me.chunkNextN(2)
			return NewLuaToken(TokenOpGe, ">=")
		}
		me.chunkNext()
		return NewLuaToken(TokenOpGt, topChar)
	case ".":
		if me.chunkStartsWith("...") {
			me.chunkNextN(3)
			return NewLuaToken(TokenVararg, "...")
		} else if me.chunkStartsWith("..") {
			me.chunkNextN(2)
			return NewLuaToken(TokenOpConcat, "..")
		}
		me.chunkNext()
		return NewLuaToken(TokenSepDot, topChar)
	case "'", "\"":
		return NewLuaToken(TokenString, me.scanShortString())
	}
	// 其他

	if me.chunkTopLikeNumber() {
		return NewLuaToken(TokenNumber, me.scanNumber())
	}

	if me.chunkTopLikeIdentifier() {
		identifier := me.scanIdentifier()
		if IsLuaKeyword(identifier) {
			return NewLuaToken(NewLuaKeyword(identifier), identifier)
		}
		return NewLuaToken(TokenIdentifier, identifier)
	}

	str := me.chunkNext()
	return NewLuaToken(TokenIdentifier, str)
}
