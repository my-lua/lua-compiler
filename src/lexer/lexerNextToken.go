package lexer

// NextToken 解析下一个单词
func (me *LuaLexer) NextToken() *LuaToken {
	// 跳过无效代码
	me.chunk.SkipInvalidCodes()
	// 如果已经没有要解析的代码，则返回EOF
	if me.chunk.IsEmpty() {
		return NewLuaToken(TokenEof, "")
	}
	// 获取顶部字符
	topChar := me.chunk.Top().CharStr()
	switch topChar {
	case ";":
		me.chunk.Next()
		return NewLuaToken(TokenSepSemi, topChar)
	case ",":
		me.chunk.Next()
		return NewLuaToken(TokenSepComma, topChar)
	case "(":
		me.chunk.Next()
		return NewLuaToken(TokenSepLParen, topChar)
	case ")":
		me.chunk.Next()
		return NewLuaToken(TokenSepRParen, topChar)
	case "[":
		if me.chunk.Top().LikeLongString() {
			return NewLuaToken(TokenString, me.chunk.ScanLongString())
		}
		me.chunk.Next()
		return NewLuaToken(TokenSepLBrack, topChar)
	case "]":
		me.chunk.Next()
		return NewLuaToken(TokenSepRBrack, topChar)
	case "{":
		me.chunk.Next()
		return NewLuaToken(TokenSepLCurly, topChar)
	case "}":
		me.chunk.Next()
		return NewLuaToken(TokenSepRCurly, topChar)
	case "+":
		me.chunk.Next()
		return NewLuaToken(TokenOpAdd, topChar)
	case "-":
		me.chunk.Next()
		return NewLuaToken(TokenOpSub, topChar)
	case "*":
		me.chunk.Next()
		return NewLuaToken(TokenOpMul, topChar)
	case "/":
		me.chunk.Next()
		return NewLuaToken(TokenOpDiv, topChar)
	case "%":
		me.chunk.Next()
		return NewLuaToken(TokenOpMod, topChar)
	case "^":
		me.chunk.Next()
		return NewLuaToken(TokenOpPow, topChar)
	case "&":
		me.chunk.Next()
		return NewLuaToken(TokenOpBAnd, topChar)
	case "|":
		me.chunk.Next()
		return NewLuaToken(TokenOpBOr, topChar)
	case "#":
		me.chunk.Next()
		return NewLuaToken(TokenOpLen, topChar)
	case ":":
		if me.chunk.Top().StartsWith("::") {
			me.chunk.NextN(2)
			return NewLuaToken(TokenSepLabel, "::")
		}
		me.chunk.Next()
		return NewLuaToken(TokenSepColon, topChar)
	case "~":
		if me.chunk.Top().StartsWith("~=") {
			me.chunk.NextN(2)
			return NewLuaToken(TokenOpNe, "~=")
		}
		me.chunk.Next()
		return NewLuaToken(TokenOpWave, topChar)
	case "=":
		if me.chunk.Top().StartsWith("==") {
			me.chunk.NextN(2)
			return NewLuaToken(TokenOpEq, "==")
		}
		me.chunk.Next()
		return NewLuaToken(TokenOpAssign, topChar)
	case "<":
		if me.chunk.Top().StartsWith("<<") {
			me.chunk.NextN(2)
			return NewLuaToken(TokenOpShl, "<<")
		} else if me.chunk.Top().StartsWith("<=") {
			me.chunk.NextN(2)
			return NewLuaToken(TokenOpLe, "<=")
		}
		me.chunk.Next()
		return NewLuaToken(TokenOpLt, topChar)
	case ">":
		if me.chunk.Top().StartsWith(">>") {
			me.chunk.NextN(2)
			return NewLuaToken(TokenOpShr, ">>")
		} else if me.chunk.Top().StartsWith(">=") {
			me.chunk.NextN(2)
			return NewLuaToken(TokenOpGe, ">=")
		}
		me.chunk.Next()
		return NewLuaToken(TokenOpGt, topChar)
	case ".":
		if me.chunk.Top().StartsWith("...") {
			me.chunk.NextN(3)
			return NewLuaToken(TokenVararg, "...")
		} else if me.chunk.Top().StartsWith("..") {
			me.chunk.NextN(2)
			return NewLuaToken(TokenOpConcat, "..")
		}
		me.chunk.Next()
		return NewLuaToken(TokenSepDot, topChar)
	case "'", "\"":
		return NewLuaToken(TokenString, me.chunk.ScanShortString())
	}
	// 其他

	if me.chunk.Top().LikeNumber() {
		return NewLuaToken(TokenNumber, me.chunk.ScanNumber())
	}

	if me.chunk.Top().LikeIdentifier() {
		identifier := me.chunk.ScanIdentifier()
		if IsLuaKeyword(identifier) {
			return NewLuaToken(NewLuaKeyword(identifier), identifier)
		}
		return NewLuaToken(TokenIdentifier, identifier)
	}

	panic("无法解析的意外字符")
}
