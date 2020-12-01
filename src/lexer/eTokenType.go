package lexer

// ETokenType 单词类型
type ETokenType int

const (
	TokenEof        ETokenType     = iota // end-of-file
	TokenVararg                           // ...
	TokenSepSemi                          // ;
	TokenSepComma                         // ,
	TokenSepDot                           // .
	TokenSepColon                         // :
	TokenSepLabel                         // ::
	TokenSepLParen                        // (
	TokenSepRParen                        // )
	TokenSepLBrack                        // [
	TokenSepRBrack                        // ]
	TokenSepLCurly                        // {
	TokenSepRCurly                        // }
	TokenOpAssign                         // =
	TokenOpMinus                          // - (sub or unm)
	TokenOpWave                           // ~ (bnot or bxor)
	TokenOpAdd                            // +
	TokenOpMul                            // *
	TokenOpDiv                            // /
	TokenOpIDiv                           // //
	TokenOpPow                            // ^
	TokenOpMod                            // %
	TokenOpBAnd                           // &
	TokenOpBOr                            // |
	TokenOpShr                            // >>
	TokenOpShl                            // <<
	TokenOpConcat                         // ..
	TokenOpLt                             // <
	TokenOpLe                             // <=
	TokenOpGt                             // >
	TokenOpGe                             // >=
	TokenOpEq                             // ==
	TokenOpNe                             // ~=
	TokenOpLen                            // #
	TokenOpAnd                            // and
	TokenOpOr                             // or
	TokenOpNot                            // not
	TokenKwBreak                          // break
	TokenKwDo                             // do
	TokenKwElse                           // else
	TokenKwElseIf                         // elseif
	TokenKwEnd                            // end
	TokenKwFalse                          // false
	TokenKwFor                            // for
	TokenKwFunction                       // function
	TokenKwGoto                           // goto
	TokenKwIf                             // if
	TokenKwIn                             // in
	TokenKwLocal                          // local
	TokenKwNil                            // nil
	TokenKwRepeat                         // repeat
	TokenKwReturn                         // return
	TokenKwThen                           // then
	TokenKwTrue                           // true
	TokenKwUntil                          // until
	TokenKwWhile                          // while
	TokenIdentifier                       // identifier
	TokenNumber                           // number literal
	TokenString                           // string literal
	TokenOpUnm      = TokenOpMinus        // unary minus
	TokenOpSub      = TokenOpMinus
	TokenOpBNot     = TokenOpWave
	TokenOpBXor     = TokenOpWave
)

var _tokenTypeNameTable = []string{
	"TokenEof",
	"TokenVararg",
	"TokenSepSemi",
	"TokenSepComma",
	"TokenSepDot",
	"TokenSepColon",
	"TokenSepLabel",
	"TokenSepLParen",
	"TokenSepRParen",
	"TokenSepLBrack",
	"TokenSepRBrack",
	"TokenSepLCurly",
	"TokenSepRCurly",
	"TokenOpAssign",
	"TokenOpMinus",
	"TokenOpWave",
	"TokenOpAdd",
	"TokenOpMul",
	"TokenOpDiv",
	"TokenOpIDiv",
	"TokenOpPow",
	"TokenOpMod",
	"TokenOpBAnd",
	"TokenOpBOr",
	"TokenOpShr",
	"TokenOpShl",
	"TokenOpConcat",
	"TokenOpLt",
	"TokenOpLe",
	"TokenOpGt",
	"TokenOpGe",
	"TokenOpEq",
	"TokenOpNe",
	"TokenOpLen",
	"TokenOpAnd",
	"TokenOpOr",
	"TokenOpNot",
	"TokenKwBreak",
	"TokenKwDo",
	"TokenKwElse",
	"TokenKwElseIf",
	"TokenKwEnd",
	"TokenKwFalse",
	"TokenKwFor",
	"TokenKwFunction",
	"TokenKwGoto",
	"TokenKwIf",
	"TokenKwIn",
	"TokenKwLocal",
	"TokenKwNil",
	"TokenKwRepeat",
	"TokenKwReturn",
	"TokenKwThen",
	"TokenKwTrue",
	"TokenKwUntil",
	"TokenKwWhile",
	"TokenIdentifier",
	"TokenNumber",
	"TokenString",
}

// getLuaTokenTypeNameTable s
func getLuaTokenTypeNameTable() []string {
	return _tokenTypeNameTable
}

// Name 获取名称
func (me ETokenType) Name() string {
	return getLuaTokenTypeNameTable()[int(me)]
}

var _keywordMap = map[string]ETokenType{
	"and":      TokenOpAnd,
	"break":    TokenKwBreak,
	"do":       TokenKwDo,
	"else":     TokenKwElse,
	"elseif":   TokenKwElseIf,
	"end":      TokenKwEnd,
	"false":    TokenKwFalse,
	"for":      TokenKwFor,
	"function": TokenKwFunction,
	"goto":     TokenKwGoto,
	"if":       TokenKwIf,
	"in":       TokenKwIn,
	"local":    TokenKwLocal,
	"nil":      TokenKwNil,
	"not":      TokenOpNot,
	"or":       TokenOpOr,
	"repeat":   TokenKwRepeat,
	"return":   TokenKwReturn,
	"then":     TokenKwThen,
	"true":     TokenKwTrue,
	"until":    TokenKwUntil,
	"while":    TokenKwWhile,
}

// getLuaKeywordMap s
func getLuaKeywordMap() map[string]ETokenType {
	return _keywordMap
}

// IsLuaKeyword 判断一个字符串是不是Lua关键字
func IsLuaKeyword(text string) bool {
	if _, found := getLuaKeywordMap()[text]; found {
		return true
	}
	return false
}

// NewLuaKeyword 根据字符串创建Lua关键字
func NewLuaKeyword(text string) ETokenType {
	if IsLuaKeyword(text) {
		return getLuaKeywordMap()[text]
	}
	panic("ETokenType NewLuaKeyword: 不是有效的Lua关键字")
}
