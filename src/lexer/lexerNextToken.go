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
	switch topChar {
	case ';':
		me.chunkNext()
		return NewToken(TokenSepSemi, string(topChar))
	case ',':
		me.chunkNext()
		return NewToken(TokenSepComma, string(topChar))
	case '(':
		me.chunkNext()
		return NewToken(TokenSepLParen, string(topChar))
	case ')':
		me.chunkNext()
		return NewToken(TokenSepRParen, string(topChar))
	case '[':
		me.chunkNext()
		return NewToken(TokenSepLBrack, string(topChar))
	case ']':
		me.chunkNext()
		return NewToken(TokenSepRBrack, string(topChar))
	case '{':
		me.chunkNext()
		return NewToken(TokenSepLCurly, string(topChar))
	case '}':
		me.chunkNext()
		return NewToken(TokenSepRCurly, string(topChar))
	}

	str := me.chunkNext()
	fmt.Print(str)
	return NewToken(TokenIdentifier, "")
}
