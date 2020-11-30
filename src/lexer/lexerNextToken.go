package lexer

import "fmt"

// NextToken 解析下一个单词
func (me *LuaLexer) NextToken() *LuaToken {
	// 跳过无效代码
	me.skipInvalidCodes()
	// 如果已经没有要解析的代码，则返回EOF
	if me.chunkIsEmpty() {
		return NewToken(TokenEof, "")
	}

	topChar := me.chunkTopChar()
	topStr := string(topChar)
	switch topChar {
	case ';':
		me.chunkNext()
		return NewToken(TokenSepSemi, topStr)
	case ',':
		me.chunkNext()
		return NewToken(TokenSepComma, topStr)
	case '(':
		me.chunkNext()
		return NewToken(TokenSepLParen, topStr)
	case ')':
		me.chunkNext()
		return NewToken(TokenSepRParen, topStr)
	case '[':
		me.chunkNext()
		return NewToken(TokenSepLBrack, topStr)
	case ']':
		me.chunkNext()
		return NewToken(TokenSepRBrack, topStr)
	case '{':
		me.chunkNext()
		return NewToken(TokenSepLCurly, topStr)
	case '}':
		me.chunkNext()
		return NewToken(TokenSepRCurly, topStr)
	case '+':
		me.chunkNext()
		return NewToken(TokenOpAdd, topStr)
	case '-':
		me.chunkNext()
		return NewToken(TokenOpSub, topStr)
	case '*':
		me.chunkNext()
		return NewToken(TokenOpMul, topStr)
	case '/':
		me.chunkNext()
		return NewToken(TokenOpDiv, topStr)
	case '%':
		me.chunkNext()
		return NewToken(TokenOpMod, topStr)
	case '^':
		me.chunkNext()
		return NewToken(TokenOpPow, topStr)
	case '&':
		me.chunkNext()
		return NewToken(TokenOpBAnd, topStr)
	case '|':
		me.chunkNext()
		return NewToken(TokenOpBOr, topStr)
	case '#':
		me.chunkNext()
		return NewToken(TokenOpLen, topStr)
	case ':':
		if me.chunkStartsWith("::") {
			me.chunkNextN(2)
			return NewToken(TokenSepLabel, "::")
		}
		me.chunkNext()
		return NewToken(TokenSepColon, topStr)
	case '~':
		if me.chunkStartsWith("~=") {
			me.chunkNextN(2)
			return NewToken(TokenOpNe, "~=")
		}
		me.chunkNext()
		return NewToken(TokenOpWave, topStr)
	case '=':
		if me.chunkStartsWith("==") {
			me.chunkNextN(2)
			return NewToken(TokenOpEq, "==")
		}
		me.chunkNext()
		return NewToken(TokenOpAssign, topStr)
	}

	str := me.chunkNext()
	fmt.Print(str)
	return NewToken(TokenIdentifier, "")
}
