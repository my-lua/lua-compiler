package lexer

// NextToken 解析下一个单词
func (me *LuaLexer) NextToken() *LuaToken {
	// 跳过无效代码
	me.chunk.SkipInvalidCodes()
	var token *LuaToken
	// 存储当前行位置，字符位置
	startLine, startChar := me.CurLine(), me.CurChar()
	// 如果已经没有要解析的代码，则返回Eof
	if me.chunk.IsEmpty() {
		token = NewLuaToken(TokenEof, "")
	} else {
		// 获取顶部字符
		topChar := me.chunk.Top().CharStr()
		switch topChar {
		case ";":
			me.chunk.Next()
			token = NewLuaToken(TokenSepSemi, topChar)
		case ",":
			me.chunk.Next()
			token = NewLuaToken(TokenSepComma, topChar)
		case "(":
			me.chunk.Next()
			token = NewLuaToken(TokenSepLParen, topChar)
		case ")":
			me.chunk.Next()
			token = NewLuaToken(TokenSepRParen, topChar)
		case "[":
			if me.chunk.Top().LikeLongString() {
				token = NewLuaToken(TokenString, me.chunk.ScanLongString())
			}
			me.chunk.Next()
			token = NewLuaToken(TokenSepLBrack, topChar)
		case "]":
			me.chunk.Next()
			token = NewLuaToken(TokenSepRBrack, topChar)
		case "{":
			me.chunk.Next()
			token = NewLuaToken(TokenSepLCurly, topChar)
		case "}":
			me.chunk.Next()
			token = NewLuaToken(TokenSepRCurly, topChar)
		case "+":
			me.chunk.Next()
			token = NewLuaToken(TokenOpAdd, topChar)
		case "-":
			me.chunk.Next()
			token = NewLuaToken(TokenOpSub, topChar)
		case "*":
			me.chunk.Next()
			token = NewLuaToken(TokenOpMul, topChar)
		case "/":
			me.chunk.Next()
			token = NewLuaToken(TokenOpDiv, topChar)
		case "%":
			me.chunk.Next()
			token = NewLuaToken(TokenOpMod, topChar)
		case "^":
			me.chunk.Next()
			token = NewLuaToken(TokenOpPow, topChar)
		case "&":
			me.chunk.Next()
			token = NewLuaToken(TokenOpBAnd, topChar)
		case "|":
			me.chunk.Next()
			token = NewLuaToken(TokenOpBOr, topChar)
		case "#":
			me.chunk.Next()
			token = NewLuaToken(TokenOpLen, topChar)
		case ":":
			if me.chunk.Top().StartsWith("::") {
				me.chunk.NextN(2)
				token = NewLuaToken(TokenSepLabel, "::")
			}
			me.chunk.Next()
			token = NewLuaToken(TokenSepColon, topChar)
		case "~":
			if me.chunk.Top().StartsWith("~=") {
				me.chunk.NextN(2)
				token = NewLuaToken(TokenOpNe, "~=")
			}
			me.chunk.Next()
			token = NewLuaToken(TokenOpWave, topChar)
		case "=":
			if me.chunk.Top().StartsWith("==") {
				me.chunk.NextN(2)
				token = NewLuaToken(TokenOpEq, "==")
			}
			me.chunk.Next()
			token = NewLuaToken(TokenOpAssign, topChar)
		case "<":
			if me.chunk.Top().StartsWith("<<") {
				me.chunk.NextN(2)
				token = NewLuaToken(TokenOpShl, "<<")
			} else if me.chunk.Top().StartsWith("<=") {
				me.chunk.NextN(2)
				token = NewLuaToken(TokenOpLe, "<=")
			}
			me.chunk.Next()
			token = NewLuaToken(TokenOpLt, topChar)
		case ">":
			if me.chunk.Top().StartsWith(">>") {
				me.chunk.NextN(2)
				token = NewLuaToken(TokenOpShr, ">>")
			} else if me.chunk.Top().StartsWith(">=") {
				me.chunk.NextN(2)
				token = NewLuaToken(TokenOpGe, ">=")
			}
			me.chunk.Next()
			token = NewLuaToken(TokenOpGt, topChar)
		case ".":
			if me.chunk.Top().StartsWith("...") {
				me.chunk.NextN(3)
				token = NewLuaToken(TokenVararg, "...")
			} else if me.chunk.Top().StartsWith("..") {
				me.chunk.NextN(2)
				token = NewLuaToken(TokenOpConcat, "..")
			}
			me.chunk.Next()
			token = NewLuaToken(TokenSepDot, topChar)
		case "'", "\"":
			token = NewLuaToken(TokenString, me.chunk.ScanShortString())
		default:
			// 其他逻辑
			if me.chunk.Top().LikeNumber() {
				token = NewLuaToken(TokenNumber, me.chunk.ScanNumber())
			} else if me.chunk.Top().LikeIdentifier() {
				identifier := me.chunk.ScanIdentifier()
				if IsLuaKeyword(identifier) {
					token = NewLuaToken(NewLuaKeyword(identifier), identifier)
				} else {
					token = NewLuaToken(TokenIdentifier, identifier)
				}
			} else {
				panic("LuaLexer NextToken: 无法解析的意外字符")
			}
		}
	}
	// 设置单词所处的行范围，字符范围
	token.SetLine(startLine, me.CurLine())
	token.SetChar(startChar, me.CurChar())
	return token
}
